package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	logPath = "logs"
	logFile = "log_%d%d%d.log"
)
var LogInstance = logrus.New()

func init() {

	currentTime := time.Now()
	logFile := fmt.Sprintf(logFile, currentTime.Year(), currentTime.Month(), currentTime.Day())
	//fmt.Println(logFile)
	// 打开文件
	logFileName, err := filepath.Abs(path.Join(logPath, logFile)) //path.Join(logPath, logFile)
	fmt.Println(logFileName)
	fileWriter, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// 设置日志输出到文件
	LogInstance.SetOutput(fileWriter)
	// 设置日志输出格式
	LogInstance.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志记录级别
	LogInstance.SetLevel(logrus.DebugLevel)
}
