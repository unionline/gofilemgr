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
)

func UpdateFiles(data string) (err error) {
	arr := []models.File{}
	json.Unmarshal([]byte(data), &arr)

	loop := len(arr)
	if loop <= 0 {
		return
	}

	for _, file := range arr {
		err = db_service.UpdateFile(&file)
		if err != nil {
			return
		}
	}
	return
}
