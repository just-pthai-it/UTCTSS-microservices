package routes

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/core/database/drivers"
	"UTCTSS-microservices/services/teacher/controllers"
	"UTCTSS-microservices/services/teacher/repositories"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes(router *gin.Engine) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")
	{
		teacherController := controllers.TeacherController{
			TeacherRepository: repositories.TeacherRepository{
				Connection: databaseConnection,
			},
		}

		apiV1RouterGroup.GET("/teachers/:id", teacherController.GetById)
		apiV1RouterGroup.POST("/teachers", teacherController.Create)
	}
}
