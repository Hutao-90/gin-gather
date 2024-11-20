package request

import (
	"bytes"
	messagePackages "debox/message"
	"debox/provider/logger"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取原始主体
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": messagePackages.ParamsError})
			return
		}
		//重新定义读取器并赋值
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		// Log request information
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"ip":           c.ClientIP(),
			"query_params": gin.H{"raw": string(body)},
		}).Info("Request received")

		// Continue processing the request
		c.Next()
	}
}

// ParseRequestParams 解析请求参数
func ParseRequestParams(c *gin.Context, params interface{}) (interface{}, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method": "ParseRequestParams",
		}).Warning(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(body, &params); err != nil {
		return nil, err
	}

	return params, nil
}
