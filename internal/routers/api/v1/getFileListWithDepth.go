/*
@Time : 2020/5/8 17:09
@Author : FB
@File : GetFileList
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gofilemgr/internal/env"
	"gofilemgr/internal/initializers/config"
	"gofilemgr/internal/services/file_service"
	"gofilemgr/pkg/util"
	"net/http"
	"strconv"
)

func GetFileListWithDepth(ctx *gin.Context) {

	var err error
	// 目录
	dir := ctx.Request.FormValue("path")
	depth := ctx.Request.FormValue("depth")

	depth_int, err := strconv.Atoi(depth)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": env.ERR_MSG_REQUEST_PARAMETER,
		})
		return
	}

	if !(depth_int >= 0 && depth_int <= 3) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": env.ERR_MSG_REQUEST_PARAMETER_DEPTH,
		})
		return
	}

	// TODO
	if dir == "" {
		dir = config.Setting.Path.FilesUploadDir
	} else {
		if !util.DirIsExist(dir) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": env.ERR_MSG_DIRECTORY_NO_FOUND,
			})
			return
		}
	}

	items, total, err := file_service.GetFileListWithDepth(dir, depth_int)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    items,
		"total":   total,
	})

	return
}
