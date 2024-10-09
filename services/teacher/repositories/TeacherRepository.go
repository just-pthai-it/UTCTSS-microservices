package repositories

import (
	"UTCTSS-microservices/core/database"
	"UTCTSS-microservices/services/teacher/models"
	_interface "UTCTSS-microservices/services/teacher/repositories/interface"
)

type TeacherRepository struct {
	_interface.ITeacherRepository
	database.Connection
}

func (teacherRepository *TeacherRepository) Create(teacher *models.Teacher) {
	teacherRepository.Connection.GormDb.Create(teacher)
}

func (teacherRepository *TeacherRepository) GetById(teacher *models.Teacher) {
	teacherRepository.Connection.GormDb.First(teacher, teacher.ID)
}
