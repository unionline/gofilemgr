/*
@Time : 2020/3/19 16:55
@Author : FB
@File : util.go
@Software: GoLand
*/
package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func PathExists(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		return
	}
	return
}

// 判定目录是否存在，不存在，则创建，否则不创建
func AutoCreateDir(path string) (ok bool) {
	err := PathExists(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			ok = false
			return
		} else {
			ok = true
			return
		}
	}

	ok = true
	return

}

func GetSystem() string {
	return runtime.GOOS
}

func IsLinuxPlatform() bool {
	return strings.EqualFold(GetSystem(), "linux")
}

func IsWindowsPlatform() bool {
	return strings.EqualFold(GetSystem(), "windows")
}

func GetPathSeparator() string {

	return string(os.PathSeparator)

}

// 在指定路径创建文件,并写入内容
func WriterFileConent(path string, data []byte) error {
	dir, _ := filepath.Split(path)
	DirIsExist(dir)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err = f.Write(data)
		return err
	}
	return err
}

// 判断目录是否存在
func DirIsExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断文件是否存在
func FileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// 新建目录
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetFileNameSuffix(filename string) string {

	pos := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename
	}
	return filename[pos+1:]
}

// windows
func ChangeFilePathR2L(filename string) string {

	var name string
	name = strings.ReplaceAll(filename, "\\\\", "/")
	name = strings.ReplaceAll(name, "\\", "/")

	return name
}

func AddFileNameSuffix(filename string, format string) string {

	pos := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename + format
	}
	if strings.Index(format, ".") == 0 {
		format = format[1:len(format)]
	}

	return filename[:pos+1] + format
}

func GetFileId(filename string) string {
	// 如果不是图片格式，根本不可能上传成功
	// ./views/resource/upload/images/abc123md5123.png -> abc123md5123
	pos := strings.LastIndex(filename, "/")
	pos2 := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename[:pos2]
	}

	return filename[pos+1 : pos2]
}

func FormatTime(t time.Time) string {

	layout := "2006-01-02 15:04:05"
	//return t.Format(time.RFC3339)
	return t.Format(layout)
}

// file

func GetFileByte(filename string) ([]byte, error) {
	byte_, err := ioutil.ReadFile(filename)
	if err != nil {
		return byte_, err
	}

	return byte_, err
}

// md5
func GetFileMd5(filename string) string {

	data_byte, err := GetFileByte(filename)
	if err != nil {
		return ""
	}

	data := MD5EncodeByte(data_byte)
	return data
}

//
func FileIsSame(filename string, filename2 string) bool {

	md5_f1 := GetFileMd5(filename)
	md5_f2 := GetFileMd5(filename2)

	if md5_f1 == md5_f2 {
		return true
	}

	return false
}
