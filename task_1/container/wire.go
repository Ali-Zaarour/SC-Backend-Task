//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"sc-backend_task.com/task_1/internal/adapter/db/repository"
	"sc-backend_task.com/task_1/internal/interface/api"
	"sc-backend_task.com/task_1/internal/usecase"
)

func WireStudentApi(db *gorm.DB) api.StudentController {
	wire.Build(repository.ProvideStudentRepository,
		usecase.ProvideFindStudentById,
		usecase.ProvideFindAllStudent,
		usecase.ProvideDeleteStudent,
		usecase.ProvidePutStudent,
		usecase.ProvidePatchStudent,
		usecase.ProvideCreateStudent,
		api.ProvideStudentController)
	return api.StudentController{}
}
