/*
@Time : 2020/4/21 16:04
@Author : FB
@File : check_image_existed
@Software: GoLand
*/
package db_service

import (
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
)

func CheckFileExisted(md5 string) (item models.File, err error) {
	item, err = Query.GetFileInfoByMd5(db.Db, md5)
	return
}
