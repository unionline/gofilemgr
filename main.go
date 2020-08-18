package main

import (
	"fmt"
	"gofilemgr/internal/initializers"
	"gofilemgr/internal/initializers/config"
	router "gofilemgr/internal/routers/server"
)

func main() {

	// 配置初始化

	initializers.Initialize()

	// 路由初始化
	r := router.Initialize()

	_ = r.Run(fmt.Sprintf(":%d", config.Setting.Port))
	//r.RunTLS(fmt.Sprintf(":%d", config.Setting.Port), config.Setting.Path.PathServerCrt, config.Setting.Path.PathServerKey)

}
