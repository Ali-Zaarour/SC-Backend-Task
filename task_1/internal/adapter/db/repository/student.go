package repository

import (
	"gorm.io/gorm"
	"sc-backend_task.com/task_1/internal/adapter/db/entity"
	"sc-backend_task.com/task_1/internal/adapter/db/mapper"
	"sc-backend_task.com/task_1/internal/domain/model"
	"sc-backend_task.com/task_1/internal/domain/repository"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
)

type StudentRepository struct {
	DB *gorm.DB
}

// ProvideStudentRepository
// struct wire provider
func ProvideStudentRepository(db *gorm.DB) repository.StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

// FindAll
// get all student with model.studentMD format
func (s StudentRepository) FindAll() ([]*model.StudentMD, error) {
	//entity of students
	var students []*entity.Student
	//handle query
	if query := s.DB.Find(&students); query.Error != nil {
		return nil, query.Error
	}
	//return mapping entity to models
	return mapper.ToStudentArrayMD(students), nil
}

// FindById
// find student data based on id
func (s StudentRepository) FindById(studentId int) (*model.StudentMD, error) {
	//entity of student
	var student *entity.Student
	//handle query
	if query := s.DB.
		Scopes(entity.ByStudentId(studentId)).
		Find(&student); query.Error != nil {
		return nil, query.Error
	}
	//return mapping entity to model
	return mapper.ToStudentMD(student), nil
}

// CreateNew
// create a new student in db
func (s StudentRepository) CreateNew(st payload.NewStudent) error {
	//entity of student
	student := &entity.Student{
		FirstName: st.FirstName,
		LastName:  st.LastName,
	}
	//handle query
	if query := s.DB.Create(&student); query.Error != nil {
		return query.Error
	}
	return nil
}

// Delete a student based on the id
func (s StudentRepository) Delete(studentId int) error {
	//entity of student
	oldStudent := &entity.Student{Id: studentId}
	//handel query
	query := s.DB.Take(&oldStudent)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected > 0 {
		if query2 := s.DB.Delete(&oldStudent); query2.Error != nil {
			return query2.Error
		}
	}
	return nil
}

// Put replace all old data in db with the new one
func (s StudentRepository) Put(student payload.Student) (*model.StudentMD, error) {
	//entity of student
	var oldStudent *entity.Student
	//handel query
	query := s.DB.Scopes(entity.ByStudentId(student.Id)).Take(&oldStudent)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected > 0 {
		oldStudent.FirstName = student.FirstName
		oldStudent.LastName = student.LastName
		if query2 := s.DB.Save(&oldStudent); query2.Error != nil {
			return nil, query2.Error
		}
	}
	return mapper.ToStudentMD(oldStudent), nil

}

// Patch replace a specific data in this entity based on id
func (s StudentRepository) Patch(student payload.Student) (*model.StudentMD, error) {
	//entity of student
	var oldStudent *entity.Student
	//handel query
	query := s.DB.Scopes(entity.ByStudentId(student.Id)).Take(&oldStudent)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected > 0 {
		entity.SetData(oldStudent, student)
		if query2 := s.DB.Save(&oldStudent); query2.Error != nil {
			return nil, query2.Error
		}
	}
	return mapper.ToStudentMD(oldStudent), nil
}
