package controllers

import (
	"UTCTSS-microservices/services/teacher/dtos"
	"UTCTSS-microservices/services/teacher/models"
	"UTCTSS-microservices/services/teacher/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TeacherController struct {
	repositories.TeacherRepository
}

// swagger:operation POST /teachers Teacher CreateTeacher
//
// # Create a teacher
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// responses:
//	 200:
//	  description: teacher
//	  schema:
//	    type: object
//	    properties:
//	      data:
//	        "$ref": "#/definitions/TeacherModel"
//	400:
//	  "$ref": "#/responses/ValidationError"

func (controller TeacherController) Create(context *gin.Context) {
	createTeacherDTO := dtos.CreateTeacherDTO{}
	if err := context.ShouldBind(&createTeacherDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher := models.NewTeacher(createTeacherDTO.ID, createTeacherDTO.Name, createTeacherDTO.IsFemale, createTeacherDTO.Birth, createTeacherDTO.UniversityTeacherDegree, createTeacherDTO.IsHeadOfDepartment, createTeacherDTO.IsHeadOfFaculty, createTeacherDTO.IsActive, createTeacherDTO.DepartmentId)
	controller.TeacherRepository.Create(&teacher)

	context.JSON(http.StatusCreated, gin.H{"data": teacher})
}

// swagger:operation GET /teachers/:id Teacher GetTeacherById
//
// # Get a teacher-form by id
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
//   - name: id
//     in: query
//     description: id of teacher need to return
//     required: true
//     type: string
// responses:
//	 201:
//	  description: teacher
//	  schema:
//	    type: object
//	    properties:
//	      data:
//	        "$ref": "#/definitions/TeacherModel"

func (controller TeacherController) GetById(context *gin.Context) {
	teacher := models.Teacher{ID: context.Param("id")}
	controller.TeacherRepository.GetById(&teacher)
	fmt.Println(teacher)
	context.JSON(http.StatusOK, gin.H{"data": teacher})
}
