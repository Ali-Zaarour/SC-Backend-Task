package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"sc-backend_task.com/task_1/container"
	"sc-backend_task.com/task_1/internal"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	db := container.GetDbConnection()

	// set route
	internal.SetRoute(e, db)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
