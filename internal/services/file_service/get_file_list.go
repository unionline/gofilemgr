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
	"strings"
)

func GetFileList(page int, limit int, file_path, file_name, file_format string) (resp []map[string]interface{}, total int, err error) {

	var offset int

	if file_path == "" || file_path == "null" {
		file_path = config.Setting.Path.FilesUploadDir
	}

	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}

	fmt.Println("page,limit,offset", page, limit, offset)
	fmt.Println("file_path", file_path)

	var file_path_normal = util.ChangeFilePathR2L(file_path)

	key_file_path_normal := file_path_normal

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
				items, err = GetFileListFromDisk(file_path, 1)
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
			//resp = make([]map[string]interface{}, len(arr))

			for i, val := range arr {
				v := val.(map[string]interface{})

				item := make(map[string]interface{})
				item = v

				// 按过滤条件赋值,格式精准匹配（txt），名称模糊匹配
				if file_name == "" && file_format == "" {
					if i >= (page-1)*limit && i < page*limit {
						//resp[i] = item
						resp = append(resp, item)
					} else if i >= page*limit {
						break
					}
				} else if file_format != "" && file_name == "" {
					if item["file_format"] == file_format {
						resp = append(resp, item)
					}

				} else if file_name != "" && file_format == "" {
					item_file_name := item["file_name"]
					if strings.Index(item_file_name.(string), file_name) > -1 {
						resp = append(resp, item)
					}

				} else if file_name != "" && file_format != "" {
					item_file_name := item["file_name"]
					if strings.Index(item_file_name.(string), file_name) > -1 && item["file_name"] == file_name {
						resp = append(resp, item)
					}
				}

			}

			return
		}
	}

	// 从数据库或者本地读起的数据
	total = len(items)
	resp = make([]map[string]interface{}, len(items))
	for i, v := range items {
		item := make(map[string]interface{})
		item["xid"] = v.XID
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
