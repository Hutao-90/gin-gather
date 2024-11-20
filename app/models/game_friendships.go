package models

import (
	"time"
)

// GameFriendships 好友关系
const (
	GameFriendshipsStatusDefault  = "0" // 初始状态
	GameFriendshipsStatusFollow   = "1"
	GameFriendshipsStatusUnFollow = "2"
)

type GameFriendships struct {
	ID        int       `json:"id" gorm:"id"`
	MemberId  int       `json:"member_id" gorm:"member_id"`   // 用户
	FriendId  int       `json:"friend_id" gorm:"friend_id"`   // 关注用户
	Status    string    `json:"status" gorm:"status"`         // 1: 关注 2:互关
	CreatedAt time.Time `json:"created_at" gorm:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"` // 更新时间
}

// TableName 表名称
func (*GameFriendships) TableName() string {
	return "game_friendships"
}
