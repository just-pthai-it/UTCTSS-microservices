package controllers

import (
	teacher_form "ESS-microservices/core/http/requests/teacher-form"
	"ESS-microservices/core/models"
	"ESS-microservices/core/repositories"
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
	input := teacher_form.CreateTeacherFormRequest{}
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher := models.NewTeacher(input.ID, input.Name, input.IsFemale, input.Birth, input.UniversityTeacherDegree, input.IsHeadOfDepartment, input.IsHeadOfFaculty, input.IsActive, input.DepartmentId)
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
