package repositories

import (
	"ESS-microservices/core/models"
	"ESS-microservices/database"
)

type TeacherRepository struct {
	database.Connection
}

func (teacherRepository *TeacherRepository) Create(teacher *models.Teacher) {
	teacherRepository.Connection.GormDb.Create(teacher)
}

func (teacherRepository *TeacherRepository) GetById(teacher *models.Teacher) {
	teacherRepository.Connection.GormDb.First(teacher, teacher.ID)
}
