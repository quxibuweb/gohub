// Package bootstrap /*
/*
* @Author: yun
* @Date:   2024/3/5 22:58
* TODO: 处理程序初始化逻辑
 */

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"quxibu/app/http/middlewares"
	"quxibu/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	// TODO: 注册全局中间键
	registerGlobalMiddleWare(router)
	// TODO: 注册 API 路由
	routes.RegisterAPIRoutes(router)
	// TODO: 处理 404 异常
	setup404Handler(router)
}

// TODO: 注册全局中间键 这里使用的是自己创建的中间键Logger() Recovery()方法 重写
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

// TODO: 处理404异常
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_no":      404,
				"error_message": "404无数据",
			})
		}
	})
}
