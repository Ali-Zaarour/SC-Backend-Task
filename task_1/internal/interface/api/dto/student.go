package dto

import "sc-backend_task.com/task_1/internal/domain/model"

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Students struct {
	NbOfStudent int       `json:"nb_of_student"`
	Students    []Student `json:"students"`
}

// ToStudentDTO
// map from model to dto
func ToStudentDTO(studentMD *model.StudentMD) *Student {
	if studentMD == nil {
		return nil
	}
	return &Student{
		Id:   studentMD.Id,
		Name: studentMD.FullName,
	}
}

// ToStudentArrayDTO
// map from []model to dto
func ToStudentArrayDTO(studentsMD []*model.StudentMD) *Students {
	if studentsMD == nil {
		return nil
	}
	var studentsDTO []Student
	for _, studentMD := range studentsMD {
		studentDTO := ToStudentDTO(studentMD)
		studentsDTO = append(studentsDTO, *studentDTO)
	}
	return &Students{
		NbOfStudent: len(studentsDTO),
		Students:    studentsDTO,
	}
}
