package router

import (
	"file/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/v1/file", handler.ListFile)
	router.POST("/v1/file", handler.UploadFile)
	router.GET("/v1/file/:fileName", handler.GetFile)
	router.DELETE("/v1/file/:fileName", handler.DeleteFile)

	return router
}
