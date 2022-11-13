package repository

import (
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

// StudentRepository
// abstract interface to  handle all student method
type StudentRepository interface {
	FindAll() ([]*model.StudentMD, error)
	FindById(studentId int) (*model.StudentMD, error)
	CreateNew(student payload.NewStudent) error
	Delete(studentId int) error
	Put(student payload.Student) (*model.StudentMD, error)
	Patch(student payload.Student) (*model.StudentMD, error)
}
