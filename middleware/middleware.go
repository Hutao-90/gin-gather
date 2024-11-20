package middleware

import (
	"debox/app/services"
	"debox/provider/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := jwt.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// 验证通过，将 claims 信息存储到 context
		member, err := services.GetMemberByWalletAddress(claims.CurrentLoginMember.WalletAddress)
		c.Set("loginMemberObject", member)
		c.Next()
	}
}
