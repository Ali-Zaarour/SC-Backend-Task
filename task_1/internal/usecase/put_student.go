package usecase

import (
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/domain/repository"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

type PutStudent struct {
	studentRepository repository.StudentRepository
}

// ProvidePutStudent
// wire provider
func ProvidePutStudent(studentRepository repository.StudentRepository) PutStudent {
	return PutStudent{
		studentRepository: studentRepository,
	}
}

func (p PutStudent) Handle(student payload.Student) (*model.StudentMD, error) {
	studentMD, err := p.studentRepository.Put(student)
	if err != nil {
		return nil, err
	}
	return studentMD, nil
}
