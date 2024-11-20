package middleware

import (
	messagePackages "debox/message"
	"debox/provider/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.LogInstance.WithFields(logrus.Fields{
					"method": "Recovery",
				}).Warning(fmt.Sprintf("Recovered from panic: %v", err))

				c.JSON(http.StatusInternalServerError, messagePackages.ServicesError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
