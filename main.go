/*
* @Author: yun
* @Date:   2024/3/5 22:40
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "Go",
		})
	})
	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusOK, "无数据")
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"error_no":      404,
				"error_message": "暂无数据",
			})
		}
	})
	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
