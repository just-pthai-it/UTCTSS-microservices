package routes

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/core/database/drivers"
	"UTCTSS-microservices/services/feedback/controllers"
	"UTCTSS-microservices/services/feedback/repositories"
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
