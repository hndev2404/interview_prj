package database

import (
	"log"

	"github.com/hndev2404/interview_beearning/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.AddressInfo{})
	Instance.AutoMigrate(&models.RelevantDetails{})
	Instance.AutoMigrate(&models.Attendance{})

	log.Println("Database Migration Completed!")
}
