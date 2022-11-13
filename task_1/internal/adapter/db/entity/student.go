package entity

import (
	"gorm.io/gorm"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

type Student struct {
	Id        int
	FirstName string
	LastName  string
}

func (s Student) TableName() string {
	return "student"
}

// ByStudentId
// return scope based on student id
func ByStudentId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id <= 0 {
			return db
		}
		return db.Where("id = ?", id)
	}
}

// SetData
// map between payload end entity on patch method
func SetData(student *Student, studentPO payload.Student) {
	if studentPO.FirstName != "" {
		student.FirstName = studentPO.FirstName
	}
	if studentPO.LastName != "" {
		student.LastName = studentPO.LastName
	}
}
