/*
@Time : 2020/4/20 14:54
@Author : FB
@File : routerLinux.go
@Software: GoLand
*/
package router

import (
	"github.com/gin-gonic/gin"
	"gofilemgr/internal/routers/api/v1"
)

func routerAPI(r *gin.Engine) {

	api := r.Group("/api")
	// 获取列表
	api.GET("/list", v1.GetFileList)
	api.GET("/list_depth", v1.GetFileListWithDepth)

	api.POST("/update", v1.UpdateFiles)
	api.POST("/delete", v1.DeleteFiles)
	api.POST("/download", v1.DownloadFiles)

	// 上传文件（可以包括图片）
	api.POST("/upload", v1.UploadFiles)
	// 上传图片
	api.POST("/upload/images/", v1.UploadImages)

	//test ok
	v12 := r.Group("/v1")
	v12.POST("/:id/xxx/", v1.UploadImages)

}
