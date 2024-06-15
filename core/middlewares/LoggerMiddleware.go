package middlewares

import (
	"ESS-microservices/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ConsoleLoggerMiddleware(context *gin.Context) {
	start := time.Now()
	context.Next()
	duration := time.Since(start)
	log.Printf("Request - Method: %s | Status: %d | Duration: %v", context.Request.Method, context.Writer.Status(), duration)
}

func FileLoggerMiddleware(context *gin.Context) {
	start := time.Now()
	context.Next()
	duration := time.Since(start)
	fileName := "request_log_" + fmt.Sprintf("%02d", start.Year()) + "-" + fmt.Sprintf("%02d", int(start.Month())) + "-" + fmt.Sprintf("%02d", start.Day()) + ".log"
	utils.WriteFile(fmt.Sprintf("Request - Method: %s | Status: %d | Duration: %v", context.Request.Method, context.Writer.Status(), duration), fileName, "./storage/logs/")
}
