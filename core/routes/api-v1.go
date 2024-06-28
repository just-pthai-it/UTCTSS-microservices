package routes

import (
	"ESS-microservices/core/controllers"
	"ESS-microservices/core/repositories"
	"ESS-microservices/database"
	"ESS-microservices/database/drivers"
	"github.com/gin-gonic/gin"
	"os"
)

// swagger:route GET /teachers/:id
func NewApiRoutes(router *gin.Engine) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")
	{
		serviceName := os.Getenv("SERVICE_NAME")
		if serviceName == "teacher" {
			teacherController := controllers.TeacherController{
				TeacherRepository: repositories.TeacherRepository{
					Connection: databaseConnection,
				},
			}

			apiV1RouterGroup.GET("/teachers/:id", teacherController.GetById)
			apiV1RouterGroup.POST("/teachers", teacherController.Create)
		}

		if serviceName == "feedback" {
			feedbackController := controllers.FeedbackController{
				Repository: repositories.FeedbackRepository{
					Connection: databaseConnection,
				},
			}

			apiV1RouterGroup.GET("/feedback/:id", feedbackController.GetById)
			apiV1RouterGroup.POST("/feedback", feedbackController.Create)
			apiV1RouterGroup.GET("/feedback", feedbackController.GetMany)
		}

	}
}
