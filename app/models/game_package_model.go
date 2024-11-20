package models

import (
	"time"
)

// GamePackage 用户包裹

type GamePackage struct {
	ID        int64     `json:"id" gorm:"id"`
	MemberId  int64     `json:"member_id" gorm:"member_id"`   // 会员ID
	Type      int8      `json:"type" gorm:"type"`             // 1:农产品 2:道具
	GoodsId   int64     `json:"goods_id" gorm:"goods_id"`     // 农产品或道具ID
	Num       int64     `json:"num" gorm:"num"`               // 数量
	Status    int8      `json:"status" gorm:"status"`         // 1:待采收 2:待出售 3:已出售 4:待使用 5:已使用 6:禁用
	CreatedAt time.Time `json:"created_at" gorm:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"` // 更新时间
}

// TableName 表名称
func (*GamePackage) TableName() string {
	return "game_package"
}
