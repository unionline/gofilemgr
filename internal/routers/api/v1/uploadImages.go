/*
@Time : 2020/4/18 14:10
@Author : FB
@File : uploadFiles
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gofilemgr/internal/services/file_service"
	"io/ioutil"
	"log"
	"net/http"
)

func UploadImages(ctx *gin.Context) {

	// 图片
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	save_path := ctx.Request.FormValue("path")

	img_path_arr := make([]string, len(files))
	for i, file := range files {
		log.Println(file.Filename)

		src, err := file.Open()
		if err != nil {
			return
		}
		defer src.Close()

		data, err := ioutil.ReadAll(src)
		if err != nil {
			return
		}

		existed, img_path, err := file_service.SaveFile(data, file, save_path)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		img_path_arr[i] = img_path

		if existed {
			continue
		}

		var dst = img_path // 上传文件到指定的路径
		ctx.SaveUploadedFile(file, dst)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
	return
}
