/*
@Time : 2020/4/16 10:15
@Author : FB
@File : file.go
@Software: GoLand
*/
package models

import (
	"gofilemgr/internal/initializers/config"
	"time"
)

type File struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	XID        string `json:"xid"`         // XID
	FileName   string `json:"file_name"`   // 原始的文件名称
	ReName     string `json:"re_name"`     // 文件重命名
	FilePath   string `json:"file_path"`   // 文件路径
	Size       string `json:"size"`        // 原始的文件大小保留两位小数
	FileFormat string `json:"file_format"` // 文件格式
	Md5        string `json:"md5"`         // MD5
	Etc        string `json:"etc"`         // 备注
	FilePerm   string `json:"file_perm"`   // 文件权限
	IsDir      bool   `json:"is_dir"`      // 是否是目录

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (File) TableName() string {
	return config.Setting.MySQL.DbPrefix + "file"
}
