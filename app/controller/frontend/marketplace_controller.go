package frontend

import (
	"debox/app/services"
	messagePackages "debox/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProductConfigAll 产品列表
func GetProductConfigAll(c *gin.Context) {

	//按照金币数倒序
	productConfigList, err := services.GetProductConfigALl("id,name,type,value")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": messagePackages.TryAgainLate})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"data":    productConfigList,
		"message": "ok",
	})
}
