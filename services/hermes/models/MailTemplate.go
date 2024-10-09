package models

import (
	"gorm.io/gorm"
	"time"
)

type MailTemplate struct {
	ID        uint           `gorm:"primaryKey" json:"id" binding:"numeric"`
	Code      string         `json:"code" binding:"required"`
	Subject   string         `json:"subject" binding:"required"`
	Content   string         `json:"content" binding:"required"`
	Layout    string         `json:"layout" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
