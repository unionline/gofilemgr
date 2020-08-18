/*
@Time : 2020/5/12 15:56
@Author : FB
@File : gofunc.go
@Software: GoLand
*/
package file_service

import (
	"gofilemgr/internal/models"
	"gofilemgr/internal/services/db_service"
	"gofilemgr/pkg/util"
	"io/ioutil"
	"log"
	"strings"
)

func updateFileMd5(files []models.File) {
	for i, v := range files {

		if v.IsDir {
			files[i].FilePath = util.ChangeFilePathR2L(files[i].FilePath)
			continue
		}

		byte_, err := ioutil.ReadFile(v.FilePath + v.FileName)
		if err != nil {
			return
		}
		if strings.LastIndex(v.Size, "GB") > -1 {
			continue
		}
		files[i].Md5 = util.MD5EncodeByte(byte_)
		files[i].XID = util.GetXID()
		files[i].FilePath = util.ChangeFilePathR2L(files[i].FilePath)
	}

	err := db_service.CreateFileBatch(&files)
	if err != nil {
		log.Println("CreateFileBatch err", err)
		return
	}
}
