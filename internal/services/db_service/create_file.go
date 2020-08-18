/*
@Time : 2020/4/17 9:27
@Author : FB
@File : deleteFile.go
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func CreateFile(item *models.File) (err error) {

	tx := db.Db.Begin()
	err = Insert.CreateFile(tx, item)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

// TODO 待修改为一次性写入
func CreateFileBatch(items *[]models.File) (err error) {

	tx := db.Db.Begin()
	for _, v := range *items {
		err = Insert.CreateFile(tx, &v)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()

	return
}
