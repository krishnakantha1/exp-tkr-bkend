package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/krishnakantha1/expenseTrackerBackend/dataaccess/mongodb"
	"github.com/krishnakantha1/expenseTrackerBackend/server"
)

func main() {
	loadEnvVariables()

}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not get environment variables.")
	}

	m := mongodb.Init("DB_URL", "DATABASE_NAME")

	server.Init(m, ":8080")

}
