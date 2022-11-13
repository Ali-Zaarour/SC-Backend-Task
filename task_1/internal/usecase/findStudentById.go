package usecase

import (
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/domain/repository"
)

type FindStudentById struct {
	StudentRepository repository.StudentRepository
}

// ProvideFindStudentById
// wire provider
func ProvideFindStudentById(studentRepository repository.StudentRepository) FindStudentById {
	return FindStudentById{
		StudentRepository: studentRepository,
	}
}

func (f FindStudentById) Handle(studentId int) (*model.StudentMD, error) {
	student, err := f.StudentRepository.FindById(studentId)
	if err != nil {
		return nil, err
	}
	return student, nil
}
