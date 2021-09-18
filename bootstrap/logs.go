package bootstrap

import (
	"github.com/gin-gonic/gin"
	"io"
	"oneQrCode/pkg/config"
	"oneQrCode/pkg/logger"
	"os"
	"strings"
)

// SetupLogs 初始化日志.
func SetupLogs() {
	filePath := config.GetString("logger.path")
	filePath = strings.TrimRight(filePath, "/")

	// 初始化 http请求日志
	logPath := filePath + "/server_http.log"
	f, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 初始化 输出日志
	fileName := "/server_out.log"
	logger.Setup(fileName, filePath)
}
