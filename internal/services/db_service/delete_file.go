/*
@Time : 2020/5/18 15:35
@Author : FB
@File : delete_file
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func DeleteFile(item *models.File) (err error) {
	tx := db.Db.Begin()
	err = Delete.DeleteByXID(tx, item)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
