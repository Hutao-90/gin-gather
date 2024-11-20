package models

import (
	"debox/provider/mysqlService"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type GameTasksModel struct {
	ID         int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string          `gorm:"type:varchar(255);not null;default:''" json:"name"`    // 任务名称
	DiamondNum int             `gorm:"default:0" json:"diamond_num"`                         // 奖励钻石数
	NFT        json.RawMessage `gorm:"type:json" json:"nft"`                                 // NFT奖励 {"nft_id":"数量"}
	Url        string          `gorm:"type:varchar(255);not null;default:''" json:"Url"`     // 任务链接
	CreatedAt  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt  time.Time       `gorm:"autoUpdateTime" json:"updated_at"`                     // 更新时间
	DeletedAt  gorm.DeletedAt  `gorm:"index" json:"deleted_at"`                              // 删除时间
}

func (GameTasksModel) TableName() string {
	return "game_task"
}

func Create(db *gorm.DB, task *GameTasksModel) error {
	return db.Create(task).Error
}

func ListAllTask() ([]GameTasksModel, error) {
	var list []GameTasksModel
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err // Return the error if the database initialization fails
	}

	// Use Find to retrieve all records
	result := db.Order("created_at desc").Find(&list)
	if result.Error != nil {
		return nil, result.Error // Return the error if the query fails
	}

	return list, nil // Return the list of tasks if successful
}
