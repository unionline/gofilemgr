/*
@Time : 2020/4/14 21:16
@Author : Justin
@Description :
@File : global.go
@Software: GoLand
*/
package env

// 全局变量
var (
	ERR_MSG_REQUEST_PARAMETER       = "请求参数不正确！"
	ERR_MSG_REQUEST_PARAMETER_DEPTH = "请求参数不正确,请勿修改路径深度参数！"

	ERR_MSG_DIRECTORY_NO_FOUND = "目录不存在"
	ERR_MSG_UNKNOWN            = "未知错误！"
)

// 全局变量，保存和请求
var (
	REQUEST_URL_RES = "/res/"
	SAVE_PAHT_RES   = "./views/resources/"
)
