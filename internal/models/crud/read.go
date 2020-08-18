/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : read.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gofilemgr/internal/models"
)

type Read struct {
}

// File
func (*Read) GetFileInfo(db *gorm.DB, xid string) (output models.File, err error) {
	err = db.First(&output, models.File{XID: xid}).Error
	return
}

func (*Read) GetFileInfoByMd5(db *gorm.DB, md5 string) (output models.File, err error) {
	err = db.Where(&models.File{Md5: md5}).First(&output).Error
	return
}

func (*Read) GetFileInfoByFileName(db *gorm.DB, fileName string) (output models.File, err error) {
	err = db.Where(&models.File{FileName: fileName}).First(&output).Error
	return
}

func (*Read) GetFileList(db *gorm.DB) (output []models.File, err error) {
	err = db.Find(&output).Error
	return
}

func (*Read) GetFileListByFilePath(db *gorm.DB, file_path string) (output []models.File, err error) {
	err = db.Where(models.File{FilePath: file_path}).Find(&output).Error
	return
}

func (*Read) GetFileRecord(db *gorm.DB, xid string) (output models.File, err error) {
	err = db.First(&output, models.File{XID: xid}).Error
	return
}
