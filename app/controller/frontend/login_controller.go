package frontend

import (
	"debox/app/services"
	messagePackages "debox/message"
	"debox/provider/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login web3钱包登录 返回token
func Login(c *gin.Context) {

	var loginParams services.LoginRequest
	if _, err := request.ParseRequestParams(c, &loginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": messagePackages.ParamsError})
		return
	}

	token, err := services.CheckWalletAddress(loginParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
