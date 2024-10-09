package _interface

import "UTCTSS-microservices/services/teacher/models"

type ITeacherRepository interface {
	ImplementRepository(model models.Teacher)
}
