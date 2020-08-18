/*
@Time : 2020/4/20 14:54
@Author : FB
@File : routerLinux.go
@Software: GoLand
*/
package router

import (
	"github.com/gin-gonic/gin"
)

func routerFile(r *gin.Engine) {

	// 定义路由
	main := r.Group("/")
	main.Static("static/", "./views/static")
	main.Static("res/", "./views/resource")
	main.Static("page/", "./views/page")
	main.StaticFile("/", "./views/index.html")
	main.StaticFile("index", "./views/index.html")
	main.StaticFile("index.html", "./views/index.html")
	main.StaticFile("favicon.ico", "./views/favicon.ico")

}
