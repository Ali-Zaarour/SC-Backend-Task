package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sc-backend_task.com/task_1/internal/interface/api/dto"
	"sc-backend_task.com/task_1/internal/interface/api/payload"
	"sc-backend_task.com/task_1/internal/usecase"
	"strconv"
)

type StudentController struct {
	FindAllStudent   usecase.FindAllStudent
	FindStudentById  usecase.FindStudentById
	CreateNewStudent usecase.CreateStudent
	DeleteStudent    usecase.DeleteStudent
	PutStudent       usecase.PutStudent
	PatchStudent     usecase.PatchStudent
}

// ProvideStudentController
// wire provider
func ProvideStudentController(
	findAllStudent usecase.FindAllStudent,
	findStudentById usecase.FindStudentById,
	createNewStudent usecase.CreateStudent,
	deleteStudent usecase.DeleteStudent,
	putStudent usecase.PutStudent,
	patchStudent usecase.PatchStudent,
) StudentController {
	return StudentController{
		FindAllStudent:   findAllStudent,
		FindStudentById:  findStudentById,
		CreateNewStudent: createNewStudent,
		DeleteStudent:    deleteStudent,
		PutStudent:       putStudent,
		PatchStudent:     patchStudent,
	}
}

func (s StudentController) FindAll(c echo.Context) error {
	students, err := s.FindAllStudent.Handle()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.FetchError})
	}
	return c.JSON(http.StatusOK, dto.ToStudentArrayDTO(students))
}

func (s StudentController) FindById(c echo.Context) error {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorWithReceivedData})
	}
	student, err := s.FindStudentById.Handle(studentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.FetchError})
	}
	return c.JSON(http.StatusOK, dto.ToStudentDTO(student))
}

func (s StudentController) Create(c echo.Context) error {
	newStudent := &payload.NewStudent{}
	if err := c.Bind(newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorWithReceivedData})
	}
	err := s.CreateNewStudent.Handel(*newStudent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorCreate})
	}
	return c.JSON(http.StatusCreated, dto.Responce{Reply: dto.Created})
}

func (s StudentController) Delete(c echo.Context) error {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorWithReceivedData})
	}
	er := s.DeleteStudent.Handle(studentId)
	if er != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorDelete})
	}
	return c.JSON(http.StatusCreated, dto.Responce{Reply: dto.Deleted})
}

func (s StudentController) Put(c echo.Context) error {
	oldStudent := &payload.Student{}
	if err := c.Bind(oldStudent); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorWithReceivedData})
	}
	newStudent, err := s.PutStudent.Handle(*oldStudent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorUpdate})
	}
	return c.JSON(http.StatusOK, dto.ToStudentDTO(newStudent))
}

func (s StudentController) Patch(c echo.Context) error {
	oldStudent := &payload.Student{}
	if err := c.Bind(oldStudent); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorWithReceivedData})
	}
	newStudent, err := s.PatchStudent.Handle(*oldStudent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Responce{Reply: dto.ErrorUpdate})
	}
	return c.JSON(http.StatusOK, dto.ToStudentDTO(newStudent))
}
