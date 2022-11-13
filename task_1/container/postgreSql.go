package container

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// GetDbConnection
// return instance of db connection
func GetDbConnection() *gorm.DB {
	Instance, dbError := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/sc-backend-task"), &gorm.Config{})
	if dbError != nil {
		log.Fatal("error configuring the database: ", dbError)
	}
	log.Println("Hey! You successfully connected to your db.")

	return Instance
}
