/*
@Time : 2020/4/21 22:59
@Author : Justin
@Description :
@File : saveFileRecord
@Software: GoLand
*/
package file_service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gofilemgr/internal/initializers/config"
	"gofilemgr/internal/models"
	"gofilemgr/internal/services/db_service"
	"gofilemgr/pkg/util"
	"mime/multipart"
)

func SaveFile(data []byte, header *multipart.FileHeader, save_path string) (existed bool, img_path string, err error) {

	item := models.File{}

	var images_dir = ""
	if save_path != "" {
		images_dir = save_path
	} else {
		images_dir = config.Setting.Path.FilesUploadDir
	}

	ok := util.AutoCreateDir(images_dir)

	if !ok {
		err = errors.New("目录无法创建")
		return
	}

	md5_str := util.MD5Encode(string(data))
	item_existed, err := db_service.CheckFileExisted(md5_str)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if item_existed.ID > 0 {
		img_path = item_existed.FilePath + item_existed.FileName
		existed = true
		return
	}

	filename := header.Filename
	size := header.Size

	// 核对名称是否重复且不同文件
	item_existed, err = db_service.CheckFileNameExisted(filename)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	var flagNameDuplicated bool
	if item_existed.ID > 0 {
		flagNameDuplicated = true
	}

	item.FileName = filename
	item.XID = util.GetXID()
	item.Size = util.ConvertHumanUnit(size)

	item.FileFormat = util.GetFileNameSuffix(filename)
	item.Md5 = md5_str
	item.FilePath = images_dir

	// 返回文件路径
	if flagNameDuplicated {
		item.ReName = filename[:len(filename)-len(item.FileFormat)-1] + "_" + util.GetRandSalt() + "." + item.FileFormat
		img_path = images_dir + item.ReName
	} else {
		img_path = images_dir + item.FileName
	}

	err = db_service.CreateFile(&item)
	if err != nil {
		return
	}

	return
}
