package _interface

import "UTCTSS-microservices/services/hermes/models"

type IMailTemplateRepository interface {
	ImplementRepository(model models.MailTemplate)
}
