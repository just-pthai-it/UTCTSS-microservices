package controllers

import (
	"UTCTSS-microservices/common"
	feedback_form "UTCTSS-microservices/core/http/requests/feedback-form"
	feedback_response "UTCTSS-microservices/core/http/responses/feedback"
	"UTCTSS-microservices/core/models"
	"UTCTSS-microservices/core/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FeedbackController struct {
	Repository repositories.FeedbackRepository
}

// swagger:operation POST /feedback Feedback CreateFeedback
//
// # Create a feedback
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// responses:
//	 200:
//	  description: created feedback
//	  schema:
//	    type: object
//	    properties:
//	      data:
//	        "$ref": "#/definitions/FeedbackModel"
//	 400:
//	  "$ref": "#/responses/ValidationError"

func (controller *FeedbackController) Create(context *gin.Context) {
	input := feedback_form.CreateFeedbackFormRequest{}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback := models.NewFeedback(models.FeedbackData{
		Title:   input.Data.Title,
		Content: input.Data.Content,
	}, input.Type, input.IsBug, input.AccountId)
	err := controller.Repository.Create(&feedback)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": feedback})
}

// swagger:operation GET /feedback/{id} Feedback GetFeedbackById
//
// # Get a feedback by id
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: id of feedback need to return
//     required: true
//     type: string
// responses:
//	 201:
//	  description: feedback
//	  schema:
//	    type: object
//	    properties:
//	      data:
//	        "$ref": "#/definitions/FeedbackModel"

func (controller *FeedbackController) GetById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	feedback := models.Feedback{ID: uint(id)}
	err := controller.Repository.GetById(&feedback)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": feedback})
}

// swagger:operation GET /feedback Feedback GetManyFeedBack
//
// # Get much feedback as pagination
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
//   - name: limit
//     in: query
//     description: maximum number of feedback need to return
//     required: false
//     type: string
//   - name: offset
//     in: query
//     description: offset of feedback need to return
//     required: false
//     type: string
//   - name: type
//     in: query
//     description: type of feedback need to return
//     required: false
//     type: string
//   - name: is_bug
//     in: query
//     description: decide list feedback need to return is bug or not
//     required: true
//     type: string
// responses:
//	 200:
//     "$ref": "#/responses/GetManyResponse"

func (controller *FeedbackController) GetMany(context *gin.Context) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(context.DefaultQuery("offset", "0"))
	typeParam := context.Query("type")
	isBugParam := context.Query("is_bug")

	conditions := map[string]any{}

	if typeParam != "" {
		conditions["type"] = typeParam
	}

	if isBugParam != "" {
		isBug, _ := strconv.ParseBool(isBugParam)
		conditions["is_bug"] = isBug
	}

	var feedback []models.Feedback
	err := controller.Repository.Paginate(&feedback, limit, offset, conditions)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	context.JSON(http.StatusOK, feedback_response.GetManyResponse{Body: feedback_response.GetManyResponseBody{Limit: limit, Offset: offset, Data: feedback}}.Body)
}
