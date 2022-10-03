package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	env "github.com/spf13/viper"
)

func ListFile(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "",
	})
}

func UploadFile(c *gin.Context) {
	filePath := env.GetString("file.path")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'file' is necessary.",
		})
		return
	}

	// fileID := uuid.NewV4().String()

	// 上傳檔案
	uploadFilePath := fmt.Sprintf("%s/%s", filePath, file.Filename)
	createDirIfNotExist(filePath) // 如果路徑不存在，自動建立目錄
	if err = c.SaveUploadedFile(file, uploadFilePath); err != nil {
		log.Error(uploadFilePath)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileName": file.Filename,
		// "size":     file.Size,
		// "mimeType": file.Header,
	})
}

func GetFile(c *gin.Context) {
	filePath := env.GetString("file.path")
	fileName := c.Param("fileName")

	downloadFilePath := fmt.Sprintf("%s/%s", filePath, fileName)

	file, err := os.Open(downloadFilePath)
	if err != nil {
		log.Error(downloadFilePath)
		panic(err)
	}
	defer file.Close()

	if _, err = io.Copy(c.Writer, file); err != nil {
		log.Error(downloadFilePath)
		panic(err)
	}
}

func DeleteFile(c *gin.Context) {
	filePath := env.GetString("file.path")
	fileName := c.Param("fileName")

	if err := os.RemoveAll(fmt.Sprintf("%s/%s", filePath, fileName)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusOK)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Error("createDirIfNotExist: ", dir)
			panic(err)
		}
	}
}
