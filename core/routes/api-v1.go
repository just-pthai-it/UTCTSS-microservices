package routes

import (
	"ESS-microservices/core/controllers"
	"ESS-microservices/core/repositories"
	"ESS-microservices/database"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /teachers/:id
func NewApiRoutes(router *gin.Engine) {
	databaseConnection := database.Connection{}
	databaseConnection.Connect()

	teacherController := controllers.TeacherController{
		TeacherRepository: repositories.TeacherRepository{
			Connection: databaseConnection,
		},
	}
	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")
	{
		apiV1RouterGroup.GET("/teachers/:id", teacherController.GetById)
		apiV1RouterGroup.POST("/teachers", teacherController.Create)
	}
}
