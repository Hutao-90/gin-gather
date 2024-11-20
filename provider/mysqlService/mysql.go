package mysqlService

import (
	sysConfig "debox/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func Init() (*gorm.DB, error) {
	//系统配置
	globalConfig := sysConfig.GetConfig()

	// 创建sql日志文件
	currentTime := time.Now()
	logFileName := fmt.Sprintf("logs/sql_%d%d%d.log", currentTime.Year(), currentTime.Month(), currentTime.Day())
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // Customize the log output
		logger.Config{
			LogLevel: logger.Info, // Set the log level (e.g., logger.Info, logger.Warn, logger.Error)
		},
	)

	dsn := globalConfig.Database.User + ":" + globalConfig.Database.Password + "@tcp(" + globalConfig.Database.Host + ":" + globalConfig.Database.Port + ")/" + globalConfig.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
