package models

import (
	customTypes "UTCTSS-microservices/custom-types"
	"gorm.io/gorm"
	"time"
)

// swagger:model TeacherModel
type Teacher struct {
	ID       string `json:"id" uri:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	IsFemale bool   `json:"is_female" binding:"boolean"`

	// swagger:strfmt date
	Birth time.Time `json:"birth" binding:"required"`

	// example: university_teacher_degree
	// extensions:
	//   x-nullable: true
	UniversityTeacherDegree customTypes.NullString `json:"university_teacher_degree" binding:"-"`
	IsHeadOfDepartment      bool                   `json:"is_head_of_department" binding:"boolean"`
	IsHeadOfFaculty         bool                   `json:"is_head_of_faculty" binding:"boolean"`
	IsActive                bool                   `json:"is_active" binding:"boolean"`
	// Extensions:
	// ---
	// x-property-value: value
	DepartmentId string `json:"department_id" gorm:"column:id_department" binding:"required"`
	// swagger:type string
	// swagger:strfmt date-time
	DeletedAt gorm.DeletedAt `json:"deleted_at" binding:"-"`
}

func (teacher *Teacher) IModelImplement() {

}

func NewTeacher(id string, name string, isFemale bool, birth string, universityTeacherDegree customTypes.NullString,
	isHeadOfDepartment bool, isHeadOfFaculty bool, isActive bool, departmentId string) Teacher {

	date, _ := time.Parse(time.DateOnly, birth)

	return Teacher{
		ID:                      id,
		Name:                    name,
		IsFemale:                isFemale,
		Birth:                   date,
		UniversityTeacherDegree: universityTeacherDegree,
		IsHeadOfDepartment:      isHeadOfDepartment,
		IsHeadOfFaculty:         isHeadOfFaculty,
		IsActive:                isActive,
		DepartmentId:            departmentId,
		DeletedAt:               gorm.DeletedAt{},
	}
}
