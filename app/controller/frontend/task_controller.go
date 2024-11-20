package frontend

import (
	"debox/app/services"
	"debox/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task represents the structure of the table
type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	DiamondNum int    `json:"diamond_num"`
	NFT        string `json:"nft"` // Using string for JSON
	Url        string `json:"url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
type TaskController struct {
	taskService *services.TaskService
}

func GetMemberTasks(c *gin.Context) {
	// 获取任务
	memberIDStr := c.PostForm("member_id")
	memberId, err := strconv.Atoi(memberIDStr)
	// 检查 member_id 是否为空
	if memberId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": messagePackages.ParamsError})
		return
	}

	// 调用 services 层的方法获取任务列表和状态
	taskList, err := services.GetTasksWithStatus(memberId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": messagePackages.TryAgainLate})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"data":    taskList,
		"message": "ok",
	})

}
