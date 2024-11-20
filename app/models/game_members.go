package models

import (
	"time"
)

// GameMembers 会员表
type GameMembers struct {
	ID             int64     `json:"id" gorm:"id"`
	DeboxUserId    int64     `json:"debox_user_id" gorm:"debox_user_id"`   // 同步debox user_id
	WalletAddress  string    `json:"wallet_address" gorm:"wallet_address"` // 钱包地址
	NickName       string    `json:"nick_name" gorm:"nick_name"`           // 账号名称
	AvatarUrl      string    `json:"avatar_url" gorm:"avatar_url"`         // 头像url
	Type           int8      `json:"type" gorm:"type"`                     // 注册类型： 1:debox 2:google 3:twitter
	NftNum         int64     `json:"nft_num" gorm:"nft_num"`               // nft数量
	Level          int8      `json:"level" gorm:"level"`                   // 用户等级
	GoldNum        int64     `json:"gold_num" gorm:"gold_num"`             // 金币数量
	DiamondNum     int64     `json:"diamond_num" gorm:"diamond_num"`       // 钻石数量
	MagicPotionNum int       `json:"magic_potion_num" gorm:"`              //魔法药水数量
	InviteCode     string    `json:"invite_code" gorm:"invite_code"`       // 钻石数量
	Status         int8      `json:"status" gorm:"status"`                 // 0:正常 1:禁用
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`         // 创建时间
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`         // 更新时间
	DeletedAt      time.Time `json:"deleted_at" gorm:"default:null"`       // 删除时间
}

// TableName 表名称
func (*GameMembers) TableName() string {
	return "game_members"
}
