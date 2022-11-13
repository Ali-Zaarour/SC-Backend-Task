package usecase

import (
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/domain/repository"
)

type FindAllStudent struct {
	StudentRepository repository.StudentRepository
}

// ProvideFindAllStudent
// wire provider
func ProvideFindAllStudent(studentRepository repository.StudentRepository) FindAllStudent {
	return FindAllStudent{
		StudentRepository: studentRepository,
	}
}

func (f FindAllStudent) Handle() ([]*model.StudentMD, error) {
	students, err := f.StudentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return students, nil
}
