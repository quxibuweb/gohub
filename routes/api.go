// Package routes /*
/*
* @Author: yun
* @Date:   2024/3/5 23:04
* 路由文件 注册路由
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "Go..",
			})
		})
	}
}
