package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/s0undy/shurl/database"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connecting to DB")
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	fmt.Println(db.Config)
}
