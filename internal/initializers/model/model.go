/*
@Time : 2020/3/12 16:50
@Author : FB
@File : model.go
@Software: GoLand
*/
package model

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func Init() {

	if !db.Db.HasTable(models.File.TableName(models.File{})) {
		db.Db.CreateTable(models.File{})
	}
	db.Db.AutoMigrate(models.File{})

}
