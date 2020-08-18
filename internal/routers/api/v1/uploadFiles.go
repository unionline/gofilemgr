/*
@Time : 2020/4/14 16:09
@Author : FB
@File : upload.go
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gofilemgr/internal/services/file_service"
	"gofilemgr/pkg/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func UploadFiles(ctx *gin.Context) {

	// 文件
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	save_path := ctx.Request.FormValue("path")
	create_dir := ctx.Request.FormValue("create_dir")
	//hide_file_show := ctx.Request.FormValue("hide_file_show")

	if util.IsLinuxPlatform() {
		//正则判断路径是否合法，只判定Linux
		//reg, err := regexp.Compile(`^(((\/home\/fanbi\/(?! )[^/\/""|\/]+)+\/?)|(\/)?)\/s*$`)
		//if err != nil {
		//	ctx.JSON(http.StatusInternalServerError, gin.H{
		//		"code":    http.StatusInternalServerError,
		//		"message": "正则不合法",
		//	})
		//	return
		//}
		//
		//ok := reg.MatchString(save_path)
		//if !ok {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"code":    http.StatusBadRequest,
		//		"message": "保存路径不合法",
		//	})
		//
		//	return
		//}

	} else if util.IsWindowsPlatform() {
		////正则判断路径是否合法
		//var err error
		//reg_r, err := regexp.Compile(`^[a-zA-Z]:(((\\(?! )[^/:*?<>\\""|\\]+)+\\?)|(\\)?)\s*$`)
		//if err != nil {
		//	return
		//}
		//reg_rr, err := regexp.Compile(`^[a-zA-Z]:(((\\\\(?! )[^/:*?<>\\""|\\\\]+)+\\\\?)|(\\\\)?)\\s*$`)
		//if err != nil {
		//	return
		//}
		//reg_l, err := regexp.Compile(`^[a-zA-Z]:(((\/(?! )[^/:*?<>\/""|\/]+)+\/?)|(\/)?)\/s*$`)
		//if err != nil {
		//	return
		//}
		//
		//var ok bool
		//
		//if reg_r.MatchString(save_path) || reg_l.MatchString(save_path) || reg_rr.MatchString(save_path) {
		//	ok = true
		//}
		//
		//if !ok {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"code":    http.StatusBadRequest,
		//		"message": "保存路径不合法",
		//	})
		//
		//	return
		//}
	}

	// 路径ChangeFilePathR2L
	save_path = util.ChangeFilePathR2L(save_path)

	img_path_arr := make([]string, len(files))
	existed_file_arr := []string{}
	for i, file := range files {
		filename := file.Filename
		log.Println(filename)

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
			existed_file_arr = append(existed_file_arr, img_path+filename)
			continue
		}

		var dst = img_path // 上传文件到指定的路径
		if create_dir == "true" {
			util.AutoCreateDir(save_path)
		} else {
			if !util.DirIsExist(save_path) {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "保存路径不存在",
				})

				return
			}
		}

		save_filename := dst

		// 是否重名的文件
		existed = util.FileIsExist(save_filename)
		if existed {

			// 判定是否同一个文件
			isSame := util.FileIsSame(save_filename, save_filename)

			if isSame {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "文件已存在",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "文件名已存在",
				})
			}

			return
		}

		ctx.SaveUploadedFile(file, dst)
	}

	if len(existed_file_arr) > 0 {

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"errcode": 1,
			"message": "文件已存在，文件名列表=" + strings.Join(existed_file_arr, ","),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"errcode": 0,
	})

	return
}
