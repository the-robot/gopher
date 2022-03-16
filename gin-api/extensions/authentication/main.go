package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"gingo/apps/db"
	errorExt "gingo/extensions/error"
	ginExt "gingo/extensions/gin"
)

type IAuthentication interface {
	Register(router *gin.RouterGroup, resource *db.Resource)
	ValidateToken(c *gin.Context)
	extractBearerToken(header string) (string, error)
	parseToken(jwtToken string) (*jwt.Token, error)
	login(s *ginExt.State) (ginExt.IResponse, errorExt.IError)
	logout(s *ginExt.State) (ginExt.IResponse, errorExt.IError)
}

type authentication struct {
	secret       string
	noAuthRoutes []string
}

func New(secret string, noAuthRoutes []string) IAuthentication {
	return &authentication{secret, noAuthRoutes}
}

func (a *authentication) Register(router *gin.RouterGroup, resource *db.Resource) {
	router.POST("/login", ginExt.RouteWrapper(resource, a.login))
	router.GET("/logout", ginExt.RouteWrapper(resource, a.logout))
}

func (a *authentication) ValidateToken(c *gin.Context) {
	if a.InNoAuthRoutes(c.Request.URL.Path) {
		c.Next()
		return
	}

	jwtToken, err := a.extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewAuthErrorResponse(err.Error(), err))
		return
	}

	token, err := a.parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewAuthErrorResponse("bad jwt token", err))
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewAuthErrorResponse("unable to parse claims", nil))
		return
	}

	// set user information
	c.Set("user", claims)
	c.Next()
}

func (a *authentication) InNoAuthRoutes(route string) bool {
	// NOTE: try to do regex no auth with query param
	for _, r := range a.noAuthRoutes {
		if r == route {
			return true
		}
	}
	return false
}

func (a *authentication) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func (a *authentication) parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(a.secret), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func (a *authentication) login(s *ginExt.State) (ginExt.IResponse, errorExt.IError) {
	loginParam := loginParameters{}
	if bindError := s.Context.ShouldBindJSON(&loginParam); bindError != nil {
		return nil, errorExt.Internal(bindError.Error(), bindError)
	}

	// create password hash
	passwordHash, hashError := loginParam.hash()
	if hashError != nil {
		return nil, hashError
	}
	fmt.Println(passwordHash)

	// NOTE: check in database for credential here

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": loginParam.Username,
		"nbf":  time.Date(2015, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenStr, signError := token.SignedString([]byte(a.secret))
	if signError != nil {
		return nil, errorExt.Internal("token sign error", signError)
	}

	return ginExt.Ok(tokenStr), nil
}

func (a *authentication) logout(_ *ginExt.State) (ginExt.IResponse, errorExt.IError) {
	return ginExt.Ok("logout successfully"), nil
}
