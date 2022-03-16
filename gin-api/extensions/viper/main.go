package viper

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	errorExt "gingo/extensions/error"
)

func setUnAuthRoutesInEnv() {
	// middleware expects it from env var, because it is a shared middleware.
	noAuth := strings.Join(viper.GetStringSlice("gin.noauth.routes")[:], ",")
	_ = os.Setenv("GIN_NOAUTH_ROUTES", noAuth)
}

func LoadViper(mode string, configPath string, absPath bool) {
	if absPath {
		configPath, _ = filepath.Abs(configPath)
	}

	// Load Configuration with viper.
	log.Infoln("Viper: running in mode " + mode)
	if mode == "test" {
		viper.SetConfigName("config.test")
	} else if mode == "prod" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("config.dev")
	}
	viper.Set("mode", mode)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	log.Infoln("[cfg]: Loading configuration config.toml from ./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Errorln(err.Error())
		panic(errorExt.Internal(err.Error(), nil))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infoln("[cfg]: Config file changed:", e.Name)
	})
	keys := viper.AllKeys()
	log.Infoln("[cfg]: Configuration loaded. Available Configuration Keys:")
	for _, v := range keys {
		log.Infoln(v)
	}
	setUnAuthRoutesInEnv()
	log.Infoln("[cfg]: Watching...")
}
