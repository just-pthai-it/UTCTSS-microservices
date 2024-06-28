package feedback_tests

import (
	"ESS-microservices/core/controllers"
	"ESS-microservices/core/repositories"
	"ESS-microservices/core/routes"
	"ESS-microservices/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestCreateFeedback(t *testing.T) {
//	feedback := models.Feedback{
//		Data: models.FeedbackData{
//			Title:   "test title",
//			Content: "test content",
//		},
//		Type:      1,
//		IsBug:     true,
//		AccountId: 1,
//	}
//
//	router := gin.Default()
//	routes.NewApiRoutes(router)
//
//	feedbackJson, _ := json.Marshal(feedback)
//
//	w := httptest.NewRecorder()
//	request, _ := http.NewRequest("POST", "/api/v1/feedback", strings.NewReader(string(feedbackJson)))
//	request.Header.Set("Content-Type", "application/json")
//
//	router.ServeHTTP(w, request)
//
//	fmt.Println(w.Body.String())
//	assert.Equal(t, 201, w.Code)
//}

//func TestCreateFeedbackValidationError(t *testing.T) {
//	feedback := models.Feedback{
//		Data: models.FeedbackData{
//			Title: "test title",
//		},
//		Type:      1,
//		IsBug:     true,
//		AccountId: 1,
//	}
//
//	router := gin.Default()
//	routes.NewApiRoutes(router)
//
//	feedbackJson, _ := json.Marshal(feedback)
//
//	w := httptest.NewRecorder()
//	request, _ := http.NewRequest("POST", "/api/v1/feedback", strings.NewReader(string(feedbackJson)))
//
//	router.ServeHTTP(w, request)
//
//	fmt.Println(w.Body.String())
//	assert.Equal(t, 400, w.Code)
//}

func TestGetFeedback(t *testing.T) {
	databaseConnection := database.Connection{}
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
	databaseConnection := database.Connection{}
	databaseConnection.Connect()

	feedbackController := controllers.FeedbackController{
		Repository: repositories.FeedbackRepository{
			Connection: databaseConnection,
		},
	}

	router := gin.Default()
	router.GET("/feedback/:id", feedbackController.GetById)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/feedback/1000", nil)
	router.ServeHTTP(w, request)

	assert.Equal(t, 404, w.Code)
}
