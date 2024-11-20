package services

import (
	"crypto/sha256"
	"debox/app/models"
	sysConfig "debox/config"
	"debox/message"
	"debox/provider/logger"
	"debox/provider/mysqlService"
	"debox/provider/web3_authenticate"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type Member struct {
	ID             int       `json:"id"`
	DeboxUserId    int       `json:"deBox_user_id"`
	WalletAddress  string    `json:"wallet_address"`
	NickName       string    `json:"nick_name"`
	AvatarUrl      string    `json:"avatar_url"`
	Type           int       `json:"type"`
	NftNum         int       `json:"nft_num"`
	Level          int       `json:"level"`            // 用户等级
	GoldNum        int       `json:"gold_num"`         // 金币数量
	DiamondNum     int       `json:"diamond_num"`      // 钻石数量
	MagicPotionNum int       `json:"magic_potion_num"` //魔法药水数量
	Status         int       `json:"status"`           // 0:正常 1:禁用
	CreatedAt      time.Time `json:"created_at"`       // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`       // 更新时间
	DeletedAt      time.Time `json:"deleted_at"`       // 更新时间
}

type AddMemberFiled struct {
	ID            int    `json:"id"`
	WalletAddress string `json:"wallet_address"`
	Status        int    `json:"status"`
}

// 定义一个数组 1=>"gold_num", 2=>"diamond_num"
var SortGoldNum = 1
var SortDiamondNum = 2
var sortNames = map[int]string{1: "gold_num", 2: "diamond_num"}

// GenerateInviteCode 生成邀请码
func GenerateInviteCode(memberID int) string {
	// Convert user ID to string
	memberIDStr := strconv.Itoa(memberID)

	// Create SHA-256 hash of the user ID
	hash := sha256.New()
	hash.Write([]byte(memberIDStr))

	// Get the hash sum and encode it to a base64 string
	inviteCode := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	// Optionally, you might want to take a substring to limit the length
	// Here, we take the first 16 characters
	if len(inviteCode) > 10 {
		inviteCode = inviteCode[:10]
	}

	return inviteCode
}

// GenerateInviteErCode 生成带链接的邀请二维码
func GenerateInviteErCode(memberID int) (string, string, error) {

	globalConfig := sysConfig.GetConfig()

	inviteCode := GenerateInviteCode(memberID)
	url := globalConfig.Site.ImagesUrl + "?invite_code=" + inviteCode

	// Generate the QR code
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GenerateInviteErCode",
			"query_params": "memberID",
		}).Warning(err.Error())
	}

	//Save the QR code to a file
	qrCodePath := fmt.Sprintf("images/%d_qrcode.png", memberID)
	err = qrCode.WriteFile(256, qrCodePath)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GenerateInviteErCode",
			"query_params": "memberID",
		}).Warning(err.Error())
	}
	return inviteCode, url, nil
}

// GetMemberByWalletAddress 检查钱包地址是否存在
func GetMemberByWalletAddress(walletAddress string) (Member, error) {
	db, err := mysqlService.Init()
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GetMemberByWalletAddress",
			"query_params": walletAddress,
		}).Warning(err.Error())
		return Member{}, err // Return the error if the database initialization fails
	}

	var memberInfo Member
	err = db.Model(&models.GameMembers{}).
		Where("wallet_address =?", walletAddress).
		First(&memberInfo).Error
	if err != nil {
		// Handle other errors
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "CheckWalletAddress",
			"query_params": "walletAddress",
		}).Warning(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return memberInfo, nil
		}
		return Member{}, err
	}

	return memberInfo, nil
}

// AddMember 通过钱包地址添加新用户
func AddMember(params LoginRequest) (memberInfo Member, err error) {
	db, err := mysqlService.Init()
	if err != nil {
		return memberInfo, err // Return the error if the database initialization fails
	}

	AddMemberRecord := models.GameMembers{
		DeboxUserId:   1,
		WalletAddress: params.WalletAddress,
		NickName:      "DeBoxUser",
		Type:          1,
		NftNum:        0,
		Level:         1,
		GoldNum:       0,
		DiamondNum:    0,
		Status:        0,
	}

	err = db.Create(&AddMemberRecord).Error
	if err != nil {
		return memberInfo, err
	}

	//生成邀请码
	GenerateMemberInviteCode(int(AddMemberRecord.ID))
	//邀请关系
	InviteCodeRegisterUser(int(AddMemberRecord.ID), params.InviteCode)

	//返回用户信息
	memberInfo.ID = int(AddMemberRecord.ID)
	memberInfo.WalletAddress = AddMemberRecord.WalletAddress
	memberInfo.Status = int(AddMemberRecord.Status)
	return memberInfo, nil
}

type LoginRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
	Message       string `json:"message" binding:"required"`
	InviteCode    string `json:"invite_code"`
}

// CheckWalletAddress 检查钱包地址是否存在并返回用户信息
func CheckWalletAddress(params LoginRequest) (string, error) {
	//用户是否已注册
	var memberInfo Member
	memberInfo, err := GetMemberByWalletAddress(params.WalletAddress)
	if err != nil {
		return "", errors.New(messagePackages.ServiceError)
	}

	// 已被禁用
	if memberInfo.Status == 1 {
		return "", errors.New(messagePackages.AccountDisable)
	}

	// TODO 是否为deBox用户
	//deBoxUserId := 1
	//用户不存在 则注册用户
	if memberInfo.DeboxUserId == 0 {
		memberInfo, err = AddMember(params)
		if err != nil {
			logger.LogInstance.WithFields(logrus.Fields{
				"method":       "CheckWalletAddress",
				"query_params": "walletAddress",
			}).Warning(err.Error())

			return "", errors.New(messagePackages.TryAgainLate)
		}
	}

	//生成token
	token, err := web3_authenticate.Authenticate(memberInfo.WalletAddress, params.Signature, params.Message)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "CheckWalletAddress",
			"query_params": params,
		}).Warning(err.Error())
		return "", err
	}

	return token, nil
}

// InviteCodeRegisterUser 邀请码注册用户
func InviteCodeRegisterUser(memberId int, inviteCode string) {
	db, err := mysqlService.Init()
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "InviteCodeRegisterUser",
			"query_params": inviteCode,
			"error":        err.Error(),
		}).Warning(err.Error())
	}

	var member models.GameMembers
	err = db.Where("invite_code =?", inviteCode).First(&member).Error
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "InviteCodeRegisterUser",
			"query_params": inviteCode,
			"error":        "不存在的邀请人",
		}).Warning(err.Error())
	}

	// 维护邀请人和被邀请人之间的关系
	var inviteCodeLogs models.GameInviteLogs
	inviteCodeLogs.FromUserId = member.ID
	inviteCodeLogs.ToUserId = int64(memberId)
	err = db.Save(&inviteCodeLogs).Error
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "InviteCodeRegisterUser",
			"query_params": inviteCodeLogs,
			"error":        err.Error(),
		}).Warning(err.Error())
	}
}

// GenerateMemberInviteCode 保存会员邀请码
func GenerateMemberInviteCode(memberId int) {
	inviteCode := GenerateInviteCode(memberId)
	db, err := mysqlService.Init()
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GenerateMemberInviteCode",
			"query_params": memberId,
			"error":        err.Error(),
		}).Warning(err.Error())
		return
	}

	var member models.GameMembers
	err = db.Where("id =?", memberId).First(&member).Error
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GenerateMemberInviteCode",
			"query_params": memberId,
			"error":        err.Error(),
		}).Warning(err.Error())
		return
	}

	member.InviteCode = inviteCode
	err = db.Save(&member).Error
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"method":       "GenerateMemberInviteCode",
			"query_params": memberId,
			"error":        err.Error(),
		}).Warning(err.Error())
	}
}

// GetMemberMapByIds 通过memberID批量查询永用户信息 返回用户列表
func GetMemberMapByIds(ids []int, sort int) (map[int]Member, error) {
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err
	}

	//sort 是否在fieldNames数组中 否则按照创建时间排序 sort则默认等于created_at字符串
	s, exists := sortNames[sort]
	if !exists {
		s = "created_at"
	}
	var members []models.GameMembers
	err = db.Where("id IN (?)", ids).Order(s + " desc").Find(&members).Error
	if err != nil {
		return nil, err
	}

	memberMap := make(map[int]Member)
	for _, member := range members {
		memberMap[int(member.ID)] = Member{
			ID:            int(member.ID),
			DeboxUserId:   int(member.DeboxUserId),
			WalletAddress: member.WalletAddress,
			NickName:      member.NickName,
			AvatarUrl:     member.AvatarUrl,
			Type:          int(member.Type),
			NftNum:        int(member.NftNum),
			Level:         int(member.Level),
			GoldNum:       int(member.GoldNum),
			DiamondNum:    int(member.DiamondNum),
		}
	}
	return memberMap, nil
}

// GetMemberById 通过ID查询用户信息
func GetMemberById(memberId int) (*Member, error) {
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err
	}

	var member models.GameMembers
	err = db.Where("id =?", memberId).First(&member).Error
	if err != nil {
		return nil, err
	}

	return &Member{
		ID:            int(member.ID),
		DeboxUserId:   int(member.DeboxUserId),
		WalletAddress: member.WalletAddress,
		NickName:      member.NickName,
		AvatarUrl:     member.AvatarUrl,
		Type:          int(member.Type),
		NftNum:        int(member.NftNum),
		Level:         int(member.Level),
		GoldNum:       int(member.GoldNum),
		DiamondNum:    int(member.DiamondNum),
		Status:        int(member.Status),
	}, nil
}
