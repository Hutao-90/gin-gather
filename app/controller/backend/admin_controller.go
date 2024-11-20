package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Admin Login",
	})
}

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "backend/dashboard.html", gin.H{
		"title": "Admin Dashboard",
	})
}
