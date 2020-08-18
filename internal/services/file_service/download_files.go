/*
@Time : 2020/5/27 9:52
@Author : FB
@File : update_files
@Software: GoLand
*/
package file_service

import (
	"encoding/json"
	"errors"
	"gofilemgr/internal/models"
	"io/ioutil"
)

func DownloadFiles(data string) ([][]byte, error) {
	arr := []models.File{}
	json.Unmarshal([]byte(data), &arr)

	loop := len(arr)
	if loop <= 0 {
		return nil, errors.New("no files selected")
	}

	data_bytes := make([][]byte, loop)
	for i, file := range arr {
		filepath := file.FilePath + file.FileName
		data_byte, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, errors.New("Read file error,file=" + filepath)
		}
		data_bytes[i] = data_byte
	}

	return data_bytes, nil
}
