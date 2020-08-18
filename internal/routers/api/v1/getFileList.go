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

func GetFileList(ctx *gin.Context) {

	var err error
	// 目录
	page := ctx.Request.FormValue("page")
	limit := ctx.Request.FormValue("limit")
	filepath := ctx.Request.FormValue("path")

	file_name := ctx.Request.FormValue("name")
	file_format := ctx.Request.FormValue("format")

	var page_int, limit_int int
	if !(page == "" || limit == "") {

		page_int, err = strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": env.ERR_MSG_REQUEST_PARAMETER,
			})
			return
		}
		limit_int, err = strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": env.ERR_MSG_REQUEST_PARAMETER,
			})
			return
		}
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": env.ERR_MSG_REQUEST_PARAMETER,
		})
		return
	}

	// TODO
	if filepath == "" {
		filepath = config.Setting.Path.FilesUploadDir
	} else {
		if !util.DirIsExist(filepath) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": env.ERR_MSG_DIRECTORY_NO_FOUND,
			})
			return
		}
	}

	items, total, err := file_service.GetFileList(page_int, limit_int, filepath, file_name, file_format)
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
