package models

import (
	"time"

	"gorm.io/gorm"
)

type GameTaskLog struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int `gorm:"not null;default:0" json:"user_id"` // 用户ID
	TaskID     int `gorm:"not null" json:"task_id"`           // 任务ID
	DiamondNum int `gorm:"default:0" json:"diamond_num"`      // 奖励钻石数
	//NFT        gorm.JSON      `json:"nft"`                                                  // NFT奖励 {"nft_id":"数量"}
	Status    int8           `gorm:"not null;default:0" json:"status"`                     // 任务状态 1：已完成
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                              // 删除时间
}

func CreateGameTaskLog(db *gorm.DB, taskLog *GameTaskLog) error {
	return db.Create(taskLog).Error
}

func GetGameTaskLogByMemberId(db *gorm.DB, memberId int) (*GameTaskLog, error) {
	var taskLog GameTaskLog
	if err := db.Where("member_id = ?", memberId).First(&taskLog).Error; err != nil {
		return nil, err
	}
	return &taskLog, nil
}
