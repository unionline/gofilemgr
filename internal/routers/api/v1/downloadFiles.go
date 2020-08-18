/*
@Time : 2020/5/27 9:42
@Author : FB
@File : downloadFiles
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gofilemgr/internal/env"
	"gofilemgr/internal/services/file_service"
	"net/http"
)

func DownloadFiles(c *gin.Context) {

	var err error

	data := c.Request.FormValue("data")

	if data == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": env.ERR_MSG_REQUEST_PARAMETER,
		})
		return
	}

	data_bytes, err := file_service.DownloadFiles(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data_bytes,
	})

	return
}
