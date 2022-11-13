package mapper

import (
	"sc-backend_task.com/task_1/internal/adapter/db/entity"
	"sc-backend_task.com/task_1/internal/domain/model"
)

// ToStudentMD
// map from entity -> to model
func ToStudentMD(student *entity.Student) *model.StudentMD {
	if student == nil {
		return nil
	}
	return &model.StudentMD{
		Id:       student.Id,
		FullName: student.FirstName + " " + student.LastName,
	}
}

// ToStudentArrayMD
// map from [] entity -> [] model
func ToStudentArrayMD(students []*entity.Student) []*model.StudentMD {
	if students == nil {
		return nil
	}
	var studentsMD []*model.StudentMD
	for _, student := range students {
		studentMD := ToStudentMD(student)
		studentsMD = append(studentsMD, studentMD)
	}
	return studentsMD
}
