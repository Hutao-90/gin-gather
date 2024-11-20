package route

import (
	"debox/app/controller/frontend"
	"debox/middleware"

	"github.com/gin-gonic/gin"
)

func SetupGroupFrontendRoutes(router *gin.Engine) {

	frontendGroup := router.Group("/api")
	// 使用恢复中间件
	frontendGroup.Use(middleware.Recovery())
	//用户登录
	frontendGroup.POST("/login", frontend.Login)
	//验权接口
	authRoute := frontendGroup.Group("", middleware.AuthMiddleware())
	{
		//签到日历详情
		authRoute.POST("/sign/logs/list", frontend.GetSignLogs)
		//签到
		authRoute.POST("/sign", frontend.SignIn)
		//获取邀请码和二维码
		authRoute.POST("/member/invite/code", frontend.GetInviteCode)
		//任务列表
		authRoute.POST("/task/list", frontend.GetMemberTasks)
		//好友列表
		authRoute.POST("/member/friend", frontend.GetFriendList)
		//关注
		authRoute.POST("/member/attention", frontend.AttentionFriend)
		//取消关注
		authRoute.POST("/member/attention/cancel", frontend.CancelAttentionFriend)
		//所有产品
		authRoute.POST("/marketplace/all", frontend.GetProductConfigAll)
	}
}
