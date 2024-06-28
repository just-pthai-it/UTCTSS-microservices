package models

import (
	"time"
)

// swagger:model FeedbackModel
type Feedback struct {
	ID        uint         `gorm:"primaryKey" json:"id" binding:"numeric"`
	Data      FeedbackData `json:"data" binding:"required" gorm:"serializer:json"`
	Type      uint         `json:"type" binding:"required"`
	IsBug     bool         `json:"is_bug" binding:"required,boolean"`
	AccountId uint         `json:"account_id" binding:"required"`
	CreatedAt time.Time    `json:"created_at"`
}

func (feedback *Feedback) IModelImplement() {

}

func (Feedback) TableName() string {
	return "feedback"
}

func NewFeedback(data FeedbackData, _type uint, isBug bool, accountId uint) Feedback {

	return Feedback{
		Data:      data,
		Type:      _type,
		IsBug:     isBug,
		AccountId: accountId,
	}
}

// swagger:model FeedbackData
type FeedbackData struct {
	// required: true
	Title string `json:"title" binding:"required"`
	// required: true
	Content string `json:"content" binding:"required"`
}
