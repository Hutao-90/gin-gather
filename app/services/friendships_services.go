package services

import (
	"debox/app/models"
	"debox/provider/logger"
	"debox/provider/mysqlService"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GameFriendships struct {
	ID        int       `json:"id"`
	MemberId  int       `json:"member_id"`  // 用户
	FriendId  int       `json:"friend_id"`  // 关注用户
	Status    string    `json:"status"`     // 1: 关注 2:互关
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at"` // 更新时间
}

type GameFriendshipsList struct {
	MemberId  int    `json:"member_id"`  // 用户
	NickName  string `json:"nick_name"`  // 昵称
	AvatarUrl string `json:"avatar_url"` // 头像
	Status    string `json:"status"`     // 1: 关注 2:互关
}

// GetMyFriendList 我的好友列表
func GetMyFriendList(memberId int, sort int) ([]GameFriendshipsList, error) {
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err
	}

	var friendList []GameFriendships
	err = db.Table("game_friendships").
		Where("member_id =?", memberId).
		Where("status>?", 0).
		Order("updated_at desc").
		Select("friend_id").
		Scan(&friendList).Error
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GetMyFriendList",
			"query_params": memberId,
			"error":        err.Error(),
		}).Warning(err.Error())
		return nil, err
	}

	if len(friendList) == 0 {
		return []GameFriendshipsList{}, nil
	}

	//读取所有好友ID并获取ID&头像&昵称
	var friendIds []int
	for _, friend := range friendList {
		friendIds = append(friendIds, friend.FriendId)
	}

	memberMap, err := GetMemberMapByIds(friendIds, sort)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GetMyFriendList",
			"query_params": memberId,
			"error":        err.Error(),
		}).Warning(err.Error())
		return nil, err
	}

	var myFriendList []GameFriendshipsList
	for _, friend := range friendList {
		member, ok := memberMap[friend.FriendId]
		if !ok {
			continue
		}
		myFriendList = append(myFriendList, GameFriendshipsList{
			MemberId:  memberId,
			Status:    friend.Status,
			NickName:  member.NickName,
			AvatarUrl: member.AvatarUrl,
		})
	}
	return myFriendList, nil
}

// AttentionFriend  关注好友
func AttentionFriend(fromMemberId int, toMemberId int) bool {

	db := GetDbConnection()
	// 用户是否存在
	var member models.GameMembers
	db.Where("id =?", toMemberId).First(&member)
	if member.ID == 0 {
		return false
	}

	// 关注用户 不存在则创建，存在则更新
	var existing models.GameFriendships
	if err := db.Where("member_id = friend_id? and ", fromMemberId, toMemberId).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//创建
			existing.MemberId = fromMemberId
			existing.FriendId = toMemberId
			if err := db.Save(&existing).Error; err != nil {
				logger.LogInstance.WithFields(logrus.Fields{
					"method":       "AttentionFriend",
					"query_params": fmt.Sprintf("fromMemberId:%d,toMemberId:%d ", fromMemberId, toMemberId),
					"error":        fmt.Sprintf("关注失败 "),
				}).Warning(err.Error())
				return false
			}
		} else {
			// Some other error occurred
			logger.LogInstance.WithFields(logrus.Fields{
				"method":       "AttentionFriend",
				"query_params": fmt.Sprintf("fromMemberId:%d,toMemberId:%d ", fromMemberId, toMemberId),
				"error":        err.Error(),
			}).Warning(err.Error())
			return false
		}
	} else {
		//已关注
		if existing.Status > models.GameFriendshipsStatusDefault {
			return true
		}
		existing.Status = models.GameFriendshipsStatusFollow
		if err := db.Save(&existing).Error; err != nil {
			logger.LogInstance.WithFields(logrus.Fields{
				"method":       "AttentionFriend",
				"query_params": fmt.Sprintf("fromMemberId:%d,toMemberId:%d ", fromMemberId, toMemberId),
				"error":        err.Error(),
			}).Warning(err.Error())
			return false
		}
	}
	return true
}

// CancelAttentionFriend 取消关注
func CancelAttentionFriend(fromMemberId int, toMemberId int) bool {

	db := GetDbConnection()
	var existing models.GameFriendships
	if err := db.Where("member_id =? and friend_id =? and status >?", fromMemberId, toMemberId, models.GameFriendshipsStatusDefault).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 未关注
			return false
		} else {
			// Some other error occurred
			logger.LogInstance.WithFields(logrus.Fields{
				"method":       "CancelAttentionFriend",
				"query_params": fmt.Sprintf("fromMemberId:%d,toMemberId:%d ", fromMemberId, toMemberId),
				"error":        err.Error(),
			}).Warning(err.Error())
			return false
		}
	}

	existing.Status = models.GameFriendshipsStatusDefault
	if err := db.Save(&existing).Error; err != nil {
		return false
	}
	return true
}
