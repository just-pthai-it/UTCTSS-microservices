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
