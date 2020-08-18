/*
@Time : 2020/4/14 22:54
@Author : Justin
@Description :
@File : delete
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gofilemgr/internal/models"
)

type Delete struct {
}

func (*Delete) DeleteByXID(db *gorm.DB, m *models.File) (err error) {
	err = db.Where(models.File{XID: m.XID}).Delete(&m).Error
	return
}
