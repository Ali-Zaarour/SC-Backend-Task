package usecase

import (
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/domain/repository"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

type PatchStudent struct {
	studentRepository repository.StudentRepository
}

// ProvidePatchStudent
// wire provider
func ProvidePatchStudent(studentRepository repository.StudentRepository) PatchStudent {
	return PatchStudent{
		studentRepository: studentRepository,
	}
}

func (p PatchStudent) Handle(student payload.Student) (*model.StudentMD, error) {
	studentMD, err := p.studentRepository.Patch(student)
	if err != nil {
		return nil, err
	}
	return studentMD, nil
}
