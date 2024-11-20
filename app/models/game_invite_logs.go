package models

import (
	"time"
)

// GameInviteLogs 邀请日志
type GameInviteLogs struct {
	ID         int64     `json:"id" gorm:"id"`
	FromUserId int64     `json:"from_user_id" gorm:"from_user_id"` // 邀请人ID
	ToUserId   int64     `json:"to_user_id" gorm:"to_user_id"`     // 被邀请人ID
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`     // 创建时间
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`     // 更新时间
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`     // 更新时间
}

// TableName 表名称
func (*GameInviteLogs) TableName() string {
	return "game_invite_logs"
}
