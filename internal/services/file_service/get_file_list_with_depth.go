/*
@Time : 2020/5/8 17:11
@Author : FB
@File : get_file_list
@Software: GoLand
*/
package file_service

import (
	"fmt"
	"github.com/go-redis/redis"
	"gofilemgr/internal/initializers/config"
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/models"
	"gofilemgr/internal/services/db_service"
	"gofilemgr/internal/services/redis_service"
	"gofilemgr/pkg/util"
	"strconv"
)

func GetFileListWithDepth(dir string, depth int) (resp []map[string]interface{}, total int, err error) {

	if dir == "" || dir == "null" {
		dir = config.Setting.Path.FilesUploadDir
	}

	fmt.Println("dir,depth", dir, depth)

	var file_path_normal = util.ChangeFilePathR2L(dir)
	depth_str := strconv.Itoa(depth)

	key_file_path_normal := file_path_normal + "_d" + depth_str

	var items []models.File
	// 首先先查redis，再查数据库，再查本地文件
	data, err := redis_service.GetRedisValueForModel(key_file_path_normal)
	if err != nil {
		// Redis未找到记录
		if err == redis.Nil {
			items, err = db_service.Query.GetFileListByFilePath(db.Db, file_path_normal)
			len_items := len(items)
			// 数据库未找到记录
			if len_items == 0 {
				items, err = GetFileListFromDisk(dir, depth)
				if err != nil {
					return
				}

				if len(items) == 0 {
					return
				}

				var files_copy = items
				go updateFileMd5(files_copy)

			}
			// 写入redis，默认24hours
			redis_service.SetRedisValue(key_file_path_normal, items)
		}
	} else {
		// 从redis缓存拿到的数据
		if arr, ok := data.([]interface{}); ok {
			total = len(arr)
			resp = make([]map[string]interface{}, len(arr))

			for i, val := range arr {
				v := val.(map[string]interface{})
				resp[i] = v
			}

			return
		}
	}

	// 从数据库或者本地读起的数据
	total = len(items)
	resp = make([]map[string]interface{}, len(items))
	for i, v := range items {
		item := make(map[string]interface{})
		item["file_name"] = v.FileName
		item["file_path"] = v.FilePath
		item["file_format"] = v.FileFormat
		item["is_dir"] = v.IsDir
		item["file_perm"] = v.FilePerm
		item["file_size"] = v.Size
		item["etc"] = v.Etc
		item["created_at"] = util.FormatTime(v.CreatedAt)

		if v.ReName != "" {
			item["file_name"] = v.ReName
		}

		resp[i] = item

	}

	return
}
