/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : create.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gofilemgr/internal/models"
)

type Insert struct {
}

func (*Insert) CreateFile(db *gorm.DB, model *models.File) (err error) {
	err = db.Create(&model).Error
	return
}
