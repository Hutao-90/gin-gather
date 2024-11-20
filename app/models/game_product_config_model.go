package models

import (
	"time"
)

// GameProductConfig 产品列表配置
const (
	GameProductConfigTypeOne   = 1
	GameProductConfigTypeTwo   = 2
	GameProductConfigTypeThree = 3
	GameProductConfigTypeFour  = 4
)

type GameProductConfig struct {
	ID        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`             // 名称
	Value     int       `json:"value" gorm:"value"`           // 值
	Type      int       `json:"type" gorm:"type"`             // 1:钻石（1U可购买钻石数量） 2:魔法药水（隐身概率） 3:邀请好友可获得钻石数 4:魔法药水价格(U或者钻石 可通过1转换)
	CreatedAt time.Time `json:"created_at" gorm:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"` // 更新时间
}

// TableName 表名称
func (*GameProductConfig) TableName() string {
	return "game_product_config"
}
