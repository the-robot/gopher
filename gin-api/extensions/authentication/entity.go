package authentication

import (
	errorExt "gingo/extensions/error"
)

type loginParameters struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *loginParameters) hash() (string, errorExt.IError) {
	return HashPassword(l.Password)
}

type AuthErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewAuthErrorResponse(message string, err error) *AuthErrorResponse {
	return &AuthErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	}
}
