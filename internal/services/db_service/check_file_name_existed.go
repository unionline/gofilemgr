/*
@Time : 2020/5/9 15:22
@Author : FB
@File : check_file_name_existed
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func CheckFileNameExisted(fileName string) (item models.File, err error) {
	item, err = Query.GetFileInfoByFileName(db.Db, fileName)
	return
}
