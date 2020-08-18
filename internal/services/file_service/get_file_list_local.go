/*
@Time : 2020/5/12 14:34
@Author : FB
@File : get_file_list_local
@Software: GoLand
*/
package file_service

import (
	"gofilemgr/internal/models"
	"gofilemgr/pkg/util"
)

func GetFileListFromDisk(dir string, depth int) (files []models.File, err error) {

	if dir != "" {
		if util.IsLinuxPlatform() {
			if dir[len(dir)-1:] != "/" {
				dir = dir + "/"
			}
		} else if util.IsWindowsPlatform() {
			if dir[len(dir)-1:] != "\\" {
				dir = dir + "\\"
			}
		}
	}

	files, err = util.WalkDirWithDepth(dir, depth)

	if err != nil {
		return
	}

	return
}
