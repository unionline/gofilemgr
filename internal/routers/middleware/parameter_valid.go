/*
@Time : 2020/4/16 15:15
@Author : FB
@File : parameter_valid.go
@Software: GoLand
*/
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validParameter(args ...string) (valid bool) {
	for _, v := range args {
		if v == "" {
			return
		}
	}
	valid = true
	return
}

// 过滤参数为空
func ValidPara(ctx *gin.Context) {

	// TODO 解析后，取不到值
	ctx.Request.ParseForm()
	val_arr := make([][]string, len(ctx.Request.PostForm))

	var i = 0
	for k, v := range ctx.Request.PostForm {

		fmt.Printf("k:%v\t", k)
		fmt.Printf("v:%v\n", v)
		val_arr[i] = v

		// TODO 当过滤表查询时，参数允许为空，不验证
		// 白名单，过滤器完善后，就不需要在每一个Controller验证参数是否为空
		{

		}

		valid := validParameter(val_arr[i]...)
		if !valid {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "参数不允许为空",
			})
			return
			// return
		}

		i++
	}

}
