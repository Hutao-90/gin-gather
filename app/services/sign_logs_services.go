package services

import (
	"debox/app/common/comfunc"
	"debox/app/models"
	"debox/provider/mysqlService"
	"time"

	"gorm.io/gorm"
)

type SignLogsService struct {
	db *gorm.DB
}

type SignLogs struct {
	Day    int `json:"day"`
	Status int `json:"status"`
}

// GetSignLogsList GetTasksWithStatus 查询用户任务及其完成状态
func GetSignLogsList(memberId int) ([]SignLogs, error) {
	year, month, days := comfunc.CurrentCalendar() // Fetch current calendar
	var signLogsStatus []SignLogs

	// Initialize the database connection
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err // Return the error if the database initialization fails
	}

	// Query the database for sign logs for the current month
	err = db.Table("game_sign_logs").
		Select("id, day, status").
		Where("member_id = ?", memberId).
		Where("year = ?", year).
		Where("month = ?", month).
		Order("id desc").
		Scan(&signLogsStatus).Error
	if err != nil {
		return nil, err // Return the error if the query fails
	}

	// Create a map to track sign status for each day
	signStatusMap := make(map[int]int)
	for _, log := range signLogsStatus {
		signStatusMap[log.Day] = log.Status // Assuming status is 1 for signed in, 0 for not signed in
	}

	// Prepare the result based on the current calendar days
	var signStatus []SignLogs
	for _, day := range days {
		status := 0 // Default to 0 (not signed in)
		if signStatus, exists := signStatusMap[day]; exists {
			status = signStatus // Update status if the user has signed in
		}
		signStatus = append(signStatus, SignLogs{Day: day, Status: status})
	}

	return signStatus, nil
}

// AddSignLog 签到
func AddSignLog(memberId, day, status int) error {
	db, err := mysqlService.Init()
	if err != nil {
		return err // Return the error if the database initialization fails
	}

	// 获取当前时间
	now := time.Now()
	year, month, day := now.Date() // Fetch current calendar

	_, err = findSignLog(memberId, year, int(month), day, 1)
	if err == nil {
		return nil // 已签到，不再插入
	}

	err = db.Create(&models.GameSignLogs{
		MemberId: memberId,
		Year:     year,
		Month:    int(month),
		Day:      day,
		Status:   status,
	}).Error

	if err != nil {
		return err // Return the error if the record insertion fails
	}

	signWeal := GameSignConfig{
		Type:       0,
		BlindId:    0,
		DiamondNum: 0,
	}
	//抽取福利
	wealType, weal := signWeal.IsOpenAndLottery()
	if weal != nil {
		if wealType == 1 {
			// 福利
			// 记录福利
			// ��除��石
			// 记录��除记录
			// 推送消息
		} else if wealType == 2 {
			// 任务
			// 记录任务
			// 推送消息
		}
	}
	return err
}

// 查询指定日期签到记录
func findSignLog(memberId int, year int, month int, day int, status int) (*SignLogs, error) {
	db, err := mysqlService.Init()
	if err != nil {
		return nil, err // Return the error if the database initialization fails
	}

	var signLog SignLogs
	err = db.Model(&models.GameSignLogs{}).
		Where("member_id =?", memberId).
		Where("year =?", year).
		Where("month =?", month).
		Where("day =?", day).
		Where("status", status).First(&models.GameSignLogs{}).Error
	if err != nil {
		return nil, err // Return the error if the record is not found or any other error occurs
	}

	return &signLog, nil
}
