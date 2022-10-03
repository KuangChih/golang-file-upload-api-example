package main

import (
	"file/api/config"
	"file/api/router"

	"github.com/spf13/viper"
)

func main() {
	config.LogInit()               // 初始化日誌系統
	config.InitApplicationConfig() // 初始化環境變數

	router := router.SetupRouter()
	port := viper.GetString("server.port")
	router.Run(":" + port)
}
