/*
@Time : 2020/5/12 14:33
@Author : FB
@File : file
@Software: GoLand
*/
package util

import (
	"gofilemgr/internal/models"
	"os"
	"path/filepath"
	"strings"
)

//获取指定目录,指定深度。
func WalkDirWithDepth(dirPth string, depth int) (files []models.File, err error) {
	files = make([]models.File, 0, 30)
	depth_dirpth := len(strings.Split(dirPth, string(os.PathSeparator)))

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}

		filenameSplit := strings.Split(filename, string(os.PathSeparator))
		depth_filename := len(filenameSplit)
		//目录默认为C:\Users\Administrator\Desktop\Temp\math\pdf，没有右斜杠，所有等depth=2 ，表示二者相减=2时，表示已经是depth为3
		if depth_filename-depth_dirpth >= depth {
			return nil
		}

		// 过滤dirPth目录
		if dirPth == filename {
			return nil
		}

		//if fi.IsDir() { // 忽略目录
		//	return nil
		//}

		item := models.File{}
		item.FileName = fi.Name()
		item.Size = ConvertHumanUnit(fi.Size())
		item.IsDir = fi.IsDir()
		item.CreatedAt = fi.ModTime()
		item.FilePerm = fi.Mode().String()

		dir, _, format, err := GetFileDirByFileNameWithPath(filename)
		if err != nil {
			return err
		}
		item.FilePath = dir
		if !fi.IsDir() {
			item.FileFormat = format
		}

		files = append(files, item)

		return nil
	})

	return files, err
}
