/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : update.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gofilemgr/internal/models"
)

type Update struct {
}

func (*Update) UpdateFile(db *gorm.DB, m *models.File) (err error) {
	err = db.Model(m).Updates(models.File{Etc: m.Etc, ReName: m.ReName, FileName: m.FileName}).Error
	return
}
