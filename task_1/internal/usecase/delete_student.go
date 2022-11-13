package usecase

import "sc-backend_task.com/task_1/internal/domain/repository"

type DeleteStudent struct {
	studentRepository repository.StudentRepository
}

// ProvideDeleteStudent
// wire provider
func ProvideDeleteStudent(studentRepository repository.StudentRepository) DeleteStudent {
	return DeleteStudent{
		studentRepository: studentRepository,
	}
}

func (d DeleteStudent) Handle(studentId int) error {
	err := d.studentRepository.Delete(studentId)
	if err != nil {
		return err
	}
	return nil
}
