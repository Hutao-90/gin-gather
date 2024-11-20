package frontend

import (
	"debox/app/services"
	messagePackages "debox/message"
	"debox/provider/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetFriendListRequest struct {
	Sort int `json:"sort"`
}

type RequestAttentionFriend struct {
	MemberId int `json:"member_id" binding:"required"`
}

// GetInviteCode 获取邀请码和二维码
func GetInviteCode(c *gin.Context) {

	loginMember, ok := services.ParseMemberInfo(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.SignatureError})
		return
	}

	// 调用 services 层的方法获取任务列表和状态
	inviteCode, qrCodePath, err := services.GenerateInviteErCode(loginMember.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messagePackages.TryAgainLate})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data": gin.H{
			"invite_code":  inviteCode,
			"qr_code_path": qrCodePath,
		},
		"message": "ok",
	})
}

// GetFriendList 获取好友列表
func GetFriendList(c *gin.Context) {
	loginMember, ok := services.ParseMemberInfo(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.SignatureError})
		return
	}

	//按照金币数倒序
	friendList, err := services.GetMyFriendList(loginMember.ID, services.SortGoldNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.TryAgainLate})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"data":    friendList,
		"message": "ok",
	})
}

// AttentionFriend 关注好友
func AttentionFriend(c *gin.Context) {
	loginMember, ok := services.ParseMemberInfo(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.SignatureError})
		return
	}

	var params RequestAttentionFriend
	if _, err := request.ParseRequestParams(c, &params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messagePackages.ParamsError})
		return
	}

	// 不能关注自己
	if loginMember.ID == params.MemberId {
		c.JSON(http.StatusOK, gin.H{"status": 1, "message": messagePackages.AttentionSelfFailure})
		return
	}

	if result := services.AttentionFriend(loginMember.ID, params.MemberId); result != true {
		c.JSON(http.StatusOK, gin.H{"error": messagePackages.AttentionFailed})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": messagePackages.AttentionSuccess})
}

// CancelAttentionFriend 取消关注
func CancelAttentionFriend(c *gin.Context) {
	loginMember, ok := services.ParseMemberInfo(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.SignatureError})
	}
	var params RequestAttentionFriend
	if _, err := request.ParseRequestParams(c, &params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messagePackages.ParamsError})
		return
	}
	if result := services.CancelAttentionFriend(loginMember.ID, params.MemberId); result != true {
		c.JSON(http.StatusOK, gin.H{"error": messagePackages.CancelAttentionFailed})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "message": messagePackages.CancelAttentionSuccess})
}
