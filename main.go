package main

import (
	database "app/config"
	"app/model"
	router "app/router"
)

// @title           Canvas Project
// @version         2.0
// @description     Api canvas project
// @host      localhost:3000
// @BasePath  /
func main() {
	db := database.MethodDB()
	db.AutoMigrate(
		model.Admin{},
		model.Teacher{},
		model.Student{},
		model.Subject{},
		model.Course{},
		model.File{},
		model.Exercise{},
		model.CourseOfStudent{},
		model.Assignment{},
		model.AttendanceStatus{},
	)
	router.App()
}
