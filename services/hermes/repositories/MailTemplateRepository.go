package repositories

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/services/hermes/models"
	_interface "UTCTSS-microservices/services/hermes/repositories/interface"
)

type MailTemplateRepository struct {
	_interface.IMailTemplateRepository
	database.Connection
}

func (repository *MailTemplateRepository) ImplementFBRepo() {
}

func (repository *MailTemplateRepository) Create(model *models.MailTemplate) error {
	return repository.Connection.GormDb.Create(model).Error
}

func (repository *MailTemplateRepository) GetById(model *models.MailTemplate) error {
	return repository.Connection.GormDb.First(model, model.ID).Error
}

func (repository *MailTemplateRepository) Paginate(MailTemplate *[]models.MailTemplate, limit int, offset int, conditions map[string]any) error {
	if conditions == nil {
		return repository.Connection.GormDb.Limit(limit).Offset(offset).Find(MailTemplate).Error
	}

	return repository.Connection.GormDb.Where(conditions).Limit(limit).Offset(offset).Find(MailTemplate).Error
}

func (repository *MailTemplateRepository) GetByCode(model *models.MailTemplate, code string) error {
	return repository.Connection.GormDb.Where(map[string]any{"code": code}).First(model).Error
}
