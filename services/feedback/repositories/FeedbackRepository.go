package repositories

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/services/feedback/models"
	_interface "UTCTSS-microservices/services/feedback/repositories/interface"
)

type FeedbackRepository struct {
	_interface.IFeedbackRepository
	database.Connection
}

func (repository *FeedbackRepository) Create(model *models.Feedback) error {
	return repository.Connection.GormDb.Create(model).Error
}

func (repository *FeedbackRepository) GetById(model *models.Feedback) error {
	return repository.Connection.GormDb.First(model, model.ID).Error
}

func (repository *FeedbackRepository) Paginate(feedback *[]models.Feedback, limit int, offset int, conditions map[string]any) error {
	if conditions == nil {
		return repository.Connection.GormDb.Limit(limit).Offset(offset).Find(feedback).Error
	}

	return repository.Connection.GormDb.Where(conditions).Limit(limit).Offset(offset).Find(feedback).Error
}
