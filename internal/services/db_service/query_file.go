/*
@Time : 2020/5/18 15:39
@Author : FB
@File : query_file
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func GetFileInfo(item models.File) (output models.File, err error) {
	tx := db.Db.Begin()
	output, err = Query.GetFileInfo(tx, item.XID)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
