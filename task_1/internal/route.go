package internal

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"sc-backend_task.com/task_1/container"
)

func SetRoute(e *echo.Echo, db *gorm.DB) {

	//Route
	Student := container.WireStudentApi(db)

	//main group
	mainGrp := e.Group("/task1")

	//student option
	stdGrp := mainGrp.Group("/student")
	stdGrp.GET("", Student.FindAll)
	stdGrp.GET("/:id", Student.FindById)
	stdGrp.POST("/new", Student.Create)
	stdGrp.DELETE("/del/:id", Student.Delete)
	stdGrp.PUT("/upd", Student.Put)
	stdGrp.PATCH("/upd", Student.Patch)
}
