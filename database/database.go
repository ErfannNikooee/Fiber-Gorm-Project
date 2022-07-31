package database

import (
	"NetFinal-Fiber_Gorm/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstanse struct {
	GOlestan *gorm.DB
}

var Database DBInstanse

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("GOlestan.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to GOlestan\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to GOlestan")
	db.Logger = logger.Default.LogMode((logger.Info))
	log.Println("Running Migrations")
	//migrations
	db.AutoMigrate(&models.Student{}, &models.Course{}, &models.GradeReport{})

	Database = DBInstanse{GOlestan: db}
}
