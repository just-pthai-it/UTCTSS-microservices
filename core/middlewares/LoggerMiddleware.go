package middlewares

import (
	"TSS-microservices/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ConsoleLoggerMiddleware(context *gin.Context) {
	start := time.Now()
	context.Next()
	end := time.Now()
	duration := time.Since(end)
	log.Printf("Request - Method: %s | Status: %d | Duration: %v | Start: %s | End: %s", context.Request.Method, context.Writer.Status(), duration, start, end)
	log.Println("===========================================================")
}

func FileLoggerMiddleware(context *gin.Context) {
	start := time.Now()
	context.Next()
	end := time.Now()
	duration := time.Since(end)
	fileName := "access_log_" + fmt.Sprintf("%02d", start.Year()) + "-" + fmt.Sprintf("%02d", int(start.Month())) + "-" + fmt.Sprintf("%02d", start.Day()) + ".log"
	utils.WriteFile(fmt.Sprintf("Request - Method: %s | Status: %d | Duration: %v | Start: %s | End: %s\n", context.Request.Method, context.Writer.Status(), duration, start, end), fileName, "./storage/logs/")
}

func GinLoggerCustomFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
