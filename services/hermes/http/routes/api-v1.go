package routes

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/core/database/drivers"
	"UTCTSS-microservices/services/hermes/controllers"
	"UTCTSS-microservices/services/hermes/repositories"
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
		hermesController := controllers.HermesController{
			Repository: repositories.MailTemplateRepository{
				Connection: databaseConnection,
			},
		}

		apiV1RouterGroup.POST("/hermes/send-mail", hermesController.SendMail)
	}
}
