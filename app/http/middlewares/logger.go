package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger HTTP 请求日志中间件
func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logFormatter,
		Output:    gin.DefaultWriter,
		SkipPaths: nil,
	})
}

// logFormatter 日志格式化.
func logFormatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[%s] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.Request.Proto,
		param.TimeStamp.Format("2006/01/02 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
