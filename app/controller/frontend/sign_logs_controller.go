package frontend

import (
	"debox/app/services"
	"debox/message"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSignLogs 签到日历详情
func GetSignLogs(c *gin.Context) {
	loginMember, ok := services.ParseMemberInfo(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.SignatureError})
	}
	// 调用 services 层的方法获取任务列表和状态
	signLogsList, err := services.GetSignLogsList(loginMember.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": messagePackages.TryAgainLate})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"data":    signLogsList,
		"message": "ok",
	})

}

// SignIn 签到
func SignIn(c *gin.Context) {
	// 获取任务
	memberIDStr := c.PostForm("member_id")
	memberId, err := strconv.Atoi(memberIDStr)

	// 检查 member_id, day, status 是否为空
	if memberId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": messagePackages.ParamsError})
		return
	}

	now := time.Now()
	_, _, day := now.Date() // Fetch current calendar

	err = services.AddSignLog(memberId, day, 1) //签到
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 0, "message": messagePackages.SignFail})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": messagePackages.SignSuccess})
}
