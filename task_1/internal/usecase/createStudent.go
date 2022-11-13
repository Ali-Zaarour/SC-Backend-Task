package usecase

import (
	"sc-backend_task.com/task_1/internal/domain/repository"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

type CreateStudent struct {
	StudentRepository repository.StudentRepository
}

// ProvideCreateStudent
// wire provider
func ProvideCreateStudent(studentRepository repository.StudentRepository) CreateStudent {
	return CreateStudent{
		StudentRepository: studentRepository,
	}
}

func (c CreateStudent) Handel(student payload.NewStudent) error {
	err := c.StudentRepository.CreateNew(student)
	if err != nil {
		return err
	}
	return nil
}
