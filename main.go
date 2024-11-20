package main

import (
	sysConfig "debox/config"
	"debox/provider/request"
	routes "debox/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

var globalConfig sysConfig.Config

func main() {

	//导入路由
	r := gin.Default()
	// 设置静态文件的目录
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	// 设置模板文件的目录
	r.LoadHTMLGlob("templates/backend/*")
	//请求日志
	r.Use(request.RequestLogger())
	// 调用路由分组的设置函数
	routes.SetupGroupFrontendRoutes(r)
	routes.SetupGroupBackendRoutes(r)

	//配置文件
	globalConfig = sysConfig.GetConfig()
	// 启动服务器并监听 8080 端口
	r.Run(fmt.Sprintf("%s:%s", globalConfig.Server.Host, globalConfig.Server.Port))

}
