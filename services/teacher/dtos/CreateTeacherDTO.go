package dtos

import customTypes "UTCTSS-microservices/custom-types"

// swagger:parameters CreateTeacher
type CreateTeacherDTO struct {
	// required: true
	ID string `json:"id" uri:"id" binding:"required"`
	// required: true
	Name string `json:"name" binding:"required"`
	// required: true
	IsFemale bool `json:"is_female" binding:"boolean"`
	// required: true
	// swagger:strfmt date
	Birth string `json:"birth" binding:"required,datetime=2006-01-02"`
	// extensions:
	//   x-nullable: true
	UniversityTeacherDegree customTypes.NullString `json:"university_teacher_degree" binding:"-"`
	// description: nullable
	IsHeadOfDepartment bool   `json:"is_head_of_department" binding:"boolean"`
	IsHeadOfFaculty    bool   `json:"is_head_of_faculty" binding:"boolean"`
	IsActive           bool   `json:"is_active" binding:"boolean"`
	DepartmentId       string `json:"department_id" gorm:"column:id_department" binding:"required"`
}
