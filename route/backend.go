package route

import (
	"debox/app/controller/backend"
	"debox/middleware"

	"github.com/gin-gonic/gin"
)

func SetupGroupBackendRoutes(router *gin.Engine) {

	backendGroup := router.Group("/bm")
	// 使用恢复中间件
	backendGroup.Use(middleware.Recovery())
	//用户登录
	backendGroup.GET("/login", backend.Login)
	//验权接口
	//authRoute := router.Group("", middleware.AuthMiddleware())
	{
		////签到日历详情
		//authRoute.POST("/sign/logs/list", controller.GetSignLogs)
		////签到
		//authRoute.POST("/sign", controller.SignIn)
		////获取邀请码和二维码
		//authRoute.POST("/member/invite/code", controller.GetInviteCode)
		////任务列表
		//authRoute.POST("/task/list", controller.GetMemberTasks)
		////好友列表
		//authRoute.POST("/member/friend", controller.GetFriendList)
		////关注
		//authRoute.POST("/member/attention", controller.AttentionFriend)
		////取消关注
		//authRoute.POST("/member/attention/cancel", controller.CancelAttentionFriend)
		////所有产品
		//authRoute.POST("/marketplace/all", controller.GetProductConfigAll)
	}
}
