package dtos

import "UTCTSS-microservices/services/feedback/models"

type CreateFeedbackDTO struct {
	// required: true
	// type: object
	Data models.FeedbackData `json:"data" binding:"required" gorm:"serializer:json"`
	// required: true
	Type uint `json:"type" binding:"required,numeric"`
	// required: true
	IsBug bool `json:"is_bug" binding:"required,boolean"`
	// required: true
	AccountId uint `json:"account_id" binding:"required,numeric"`
}

// swagger:parameters CreateFeedback
type Payload struct {
	// required: true
	// in: body
	// type: object
	Body CreateFeedbackDTO
}
