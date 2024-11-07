package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/krishnakantha1/expenseTrackerBackend/dataaccess/mongodb"
	"github.com/krishnakantha1/expenseTrackerBackend/server"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
)

func main() {
	loadEnvVariables()

}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not get environment variables: ", err)
	}

	m := mongodb.Init("DB_URL", "DATABASE_NAME")

	port, err := utils.GetEnv("PORT")
	if len(port) == 0 || err != nil {
		port = ":8080"
	}

	server.Init(m, port)

}
