package router

import (
	"github.com/gin-gonic/gin"
)

// Initialize ...
func Initialize() *gin.Engine {

	var r = gin.Default()

	// 校验参数
	routerFile(r)

	//r.Use(middleware.ValidPara)
	routerAPI(r)

	return r
}
