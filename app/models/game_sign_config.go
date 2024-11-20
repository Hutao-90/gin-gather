package models

import (
	"time"
)

// GameSignConfig 签到配置
type GameSignConfig struct {
	ID          int       `json:"id" gorm:"id"`
	Date        time.Time `json:"date" gorm:"date"`               // 日期
	Type        int       `json:"type" gorm:"type"`               // 1:普通盲盒 2:大额盲盒
	BlindId     int       `json:"blind_id" gorm:"blind_id"`       // 盲盒ID
	DiamondNum  int       `json:"diamond_num" gorm:"diamond_num"` // 钻石数量
	Probability int       `json:"probability" gorm:"probability"` // 抽中概率
	Status      int       `json:"status" gorm:"status"`           // 0:关闭 1:开放
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`   // 更新时间
	DeletedAt   time.Time `json:"deleted_at" gorm:"deleted_at"`   // 更新时间
}

// TableName 表名称
func (*GameSignConfig) TableName() string {
	return "game_sign_config"
}
