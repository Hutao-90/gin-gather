package models

import (
	"time"

	"gorm.io/gorm"
)

type GameSignLogs struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	MemberId  int            `gorm:"not null;default:0;index" json:"member_id"`            // 同步 debox user_id
	Year      int            `gorm:"not null;default:0" json:"year"`                       // 年
	Month     int            `gorm:"not null;default:0" json:"month"`                      // 月
	Day       int            `gorm:"not null;default:0" json:"day"`                        // 日
	Status    int            `gorm:"not null;default:0" json:"status"`                     // 状态
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                              // 删除时间
}
