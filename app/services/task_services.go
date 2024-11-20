package services

import (
	"debox/provider/mysqlService"

	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

type TaskStatus struct {
	ID     int    `json:"task_id"`
	Name   string `json:"task_name"`
	Url    string `json:"url"`
	Status int    `json:"status"` // 1: 已完成, 0: 未完成
}

// GetTasksWithStatus 查询用户任务及其完成状态
func GetTasksWithStatus(memberId int) ([]TaskStatus, error) {
	var tasksWithStatus []TaskStatus

	db, err := mysqlService.Init()
	if err != nil {
		return nil, err // Return the error if the database initialization fails
	}

	err = db.Table("game_tasks").
		Select("game_tasks.id, game_tasks.name, game_tasks.url, IF(game_task_logs.id IS NOT NULL, 1, 0) AS status").
		// Left join task_logs to include tasks even if there's no log entry
		Joins("LEFT JOIN game_task_logs ON game_tasks.id = game_task_logs.task_id AND game_task_logs.member_id = ?", memberId).
		Order("game_tasks.id desc").
		Scan(&tasksWithStatus).Error

	return tasksWithStatus, err
}
