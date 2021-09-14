package bootstrap

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// SetupLogs 初始化日志.
func SetupLogs() {
	f, _ := os.Create("logs/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
