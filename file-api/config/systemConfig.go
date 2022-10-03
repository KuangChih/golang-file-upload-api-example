package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitApplicationConfig() error {

	if err := InitProfileConfig("application"); err != nil {
		return err
	}

	profile := viper.GetString("profiles.active")

	err := InitProfileConfig(profile)

	return err
}

func InitProfileConfig(envType string) error {
	log.Info("Using Profile: ", envType)

	// 新增key replacer
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(envType + ".yml")
	viper.AddConfigPath("./configYml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
		return err
	}

	// 取得系統環境變數，複寫yml中的變數
	viper.AutomaticEnv()

	return nil
}

func GetEnv(argName string) interface{} {
	return viper.Get(argName)
}
