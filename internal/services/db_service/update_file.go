/*
@Time : 2020/5/18 15:34
@Author : FB
@File : update_file
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func UpdateFile(item *models.File) (err error) {

	tx := db.Db.Begin()
	err = Update.UpdateFile(tx, item)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}
