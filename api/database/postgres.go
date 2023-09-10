package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	//Table-Model
	type Shurl struct {
		gorm.Model
		ID   uint64 `gorm:"primaryKey;autoIncrement:false"`
		Lurl string
		Surl uint64
		Uuid string
	}

	//Define Data source
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
		os.Getenv("TimeZone"),
	)

	//Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connec to database")
	} else {
		log.Println("Succesfully connected to DB")
	}

	//Check if Shurl table exsists in db, if not create it
	table := db.Migrator().HasTable(&Shurl{})
	if !table {
		log.Println("Database not initialized, creating tables")
		err := db.AutoMigrate(&Shurl{})
		if err != nil {
			log.Println("Succesfully initialized the DB")
		} else {
			log.Println("Table has been created")
		}
	} else {
		log.Println("Database already initialized, skipping initialization")
	}
	return db, nil
}
