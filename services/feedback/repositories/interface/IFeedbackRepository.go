package _interface

import "UTCTSS-microservices/services/feedback/models"

type IFeedbackRepository interface {
	ImplementRepository(model models.Feedback)
}
