package feedback_tests

import (
	"UTCTSS-microservices/core/controllers"
	"UTCTSS-microservices/core/models"
	"UTCTSS-microservices/core/repositories"
	"UTCTSS-microservices/core/routes"
	"UTCTSS-microservices/database"
	"UTCTSS-microservices/database/drivers"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateFeedback(t *testing.T) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	feedbackController := controllers.FeedbackController{
		Repository: repositories.FeedbackRepository{
			Connection: databaseConnection,
		},
	}

	router := gin.Default()
	router.POST("/api/v1/feedback", feedbackController.Create)

	feedback := models.Feedback{
		Data: models.FeedbackData{
			Title:   "test title",
			Content: "test content",
		},
		Type:      1,
		IsBug:     true,
		AccountId: 1,
	}

	feedbackJson, _ := json.Marshal(feedback)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/feedback", strings.NewReader(string(feedbackJson)))
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, request)

	assert.Equal(t, 201, w.Code)
}

func TestCreateFeedbackValidationError(t *testing.T) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	feedbackController := controllers.FeedbackController{
		Repository: repositories.FeedbackRepository{
			Connection: databaseConnection,
		},
	}

	router := gin.Default()
	router.POST("/api/v1/feedback", feedbackController.Create)

	feedback := models.Feedback{
		Data: models.FeedbackData{
			Title: "test title",
		},
		Type:      1,
		IsBug:     true,
		AccountId: 1,
	}

	feedbackJson, _ := json.Marshal(feedback)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/feedback", strings.NewReader(string(feedbackJson)))

	router.ServeHTTP(w, request)

	assert.Equal(t, 400, w.Code)
}

func TestGetFeedback(t *testing.T) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	feedbackController := controllers.FeedbackController{
		Repository: repositories.FeedbackRepository{
			Connection: databaseConnection,
		},
	}

	router := gin.Default()
	router.GET("/feedback/:id", feedbackController.GetById)

	routes.NewApiRoutes(router)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/feedback/1", nil)
	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
}

func TestGetFeedbackNotFound(t *testing.T) {
	databaseConnection := database.Connection{
		Driver: &drivers.PostgresDriver{},
	}
	databaseConnection.Connect()

	feedbackController := controllers.FeedbackController{
		Repository: repositories.FeedbackRepository{
			Connection: databaseConnection,
		},
	}

	router := gin.Default()
	router.GET("/feedback/:id", feedbackController.GetById)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/feedback/10000", nil)
	router.ServeHTTP(w, request)

	assert.Equal(t, 404, w.Code)
}
