package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ERR_NO_FOUND_R      = errors.New("没有找到右斜杠")
	ERR_NO_FOUND_FORMAT = errors.New("没有找到后缀名")
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func ConvertHumanUnit(num_int64 int64) string {

	return ByteSize(float64(num_int64)).String()

}

func Str2Float(str string) (f float64) {
	f, _ = strconv.ParseFloat(str, 64)
	return
}

func Float2Str(f float64, prec int) (str string) {
	str = strconv.FormatFloat(f, 'f', prec, 64)
	return
}

func Str2FloatPrec(strFloat string, prec int) (str string) {

	if prec == -1 {
		return Float2Str(Str2Float(strFloat), -1)
	}

	pos := strings.Index(strFloat, ".")
	pos2 := pos + prec + 1
	len_str := len(strFloat)

	if pos == -1 {
		str = strFloat
	} else if pos2 <= len_str {
		str = strFloat[:pos+prec+1]
	} else {
		str = strFloat[:len_str]
	}

	return
}

func GetFileDirByFileNameWithPath(name string) (dir string, filename string, format string, err error) {
	pos := strings.LastIndex(name, string(os.PathSeparator))
	if pos <= 0 {
		err = ERR_NO_FOUND_R
		return
	}

	pos_suffix := strings.LastIndex(name, ".")
	if pos <= 0 {
		err = ERR_NO_FOUND_FORMAT
		return
	}

	//// /fanbi/abc/img001.png
	// dir,filename,format=>/fanbi/abc/,img001.png,png
	dir = name[:pos+1]
	filename = name[pos+1:]
	format = name[pos_suffix+1:]

	return
}
