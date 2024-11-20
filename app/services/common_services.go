package services

import (
	"debox/provider/logger"
	"debox/provider/mysqlService"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ParseMemberInfo 解析登录用户信息
func ParseMemberInfo(c *gin.Context) (Member, bool) {
	loginMemberObject, exists := c.Get("loginMemberObject")
	if !exists {
		logger.LogInstance.WithFields(logrus.Fields{
			"method": "parseMemberInfo",
		}).Warning(exists)
		return Member{}, exists
	}

	loginMember, ok := loginMemberObject.(Member)
	if !ok {
		logger.LogInstance.WithFields(logrus.Fields{
			"method": "parseMemberInfo",
		}).Warning(ok)
		return Member{}, ok
	}

	return loginMember, true
}

// GetDbConnection 获取数据库操作实例
func GetDbConnection() *gorm.DB {
	db, err := mysqlService.Init()
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method": "GetDbConnection",
			"error":  err.Error(),
		}).Error(err)
		panic(err)
	}
	return db
}
