/*
@Time : 2020/5/27 9:52
@Author : FB
@File : update_files
@Software: GoLand
*/
package file_service

import (
	"encoding/json"
	"gofilemgr/internal/models"
	"gofilemgr/internal/services/db_service"
	"os"
)

func DeleteFiles(data string) (err error) {
	arr := []models.File{}
	json.Unmarshal([]byte(data), &arr)

	loop := len(arr)
	if loop <= 0 {
		return
	}

	for _, file := range arr {
		// 查询记录
		var item models.File
		item, err = db_service.GetFileInfo(file)
		if err != nil {
			return
		}

		//删除数据库文件
		err = db_service.DeleteFile(&file)
		if err != nil {
			return
		}

		//删除本地文件
		// 删除本地文件
		fileName := item.FilePath + item.FileName
		err = os.Remove(fileName)
		if err != nil {
			return
		}
	}

	return
}
