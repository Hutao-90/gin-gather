package services

import (
	"math/rand"
	"time"
)

// GameSignConfig 签到相关
type GameSignConfig struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`        // 日期
	Type        int       `json:"type"`        // 1:普通盲盒 2:大额盲盒
	BlindId     int       `json:"blind_id"`    // 盲盒ID
	DiamondNum  int       `json:"diamond_num"` // 钻石数量
	Probability int       `json:"probability"` // 抽中概率
	Status      int       `json:"status"`      // 0:关闭 1:开放
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
	DeletedAt   time.Time `json:"deleted_at"`  // 更新时间
}

// IsOpenAndLottery 按照日期判读当天是否开放抽奖且当前类型是抽取普通盲盒还是大额盲盒
// 通过一定的概率来判读是否中奖，如果抽中大额盲盒则返回类型和盲盒盲盒ID，如果抽中普通盲盒则返回类型和钻石数量
func (gsc *GameSignConfig) IsOpenAndLottery() (int, interface{}) {
	if gsc.Status == 0 {
		return 0, nil
	}

	now := time.Now()
	if gsc.Date.Year() != now.Year() || gsc.Date.Month() != now.Month() || gsc.Date.Day() != now.Day() {
		return 0, nil
	}

	//大额盲盒
	randNum := rand.Intn(100) + 1
	if gsc.Type == 2 {
		if randNum >= gsc.Probability {
			return gsc.Type, gsc.BlindId
		} else {
			return gsc.Type, nil
		}
	}

	//普通盲盒
	if randNum >= gsc.Probability {
		return gsc.Type, gsc.DiamondNum
	} else {
		return gsc.Type, nil
	}

}
