package main

import (
	"log"

	"oms-test/api"
	"oms-test/database"
)

func main() {
	// Connect to the database
	database.Connect()

	// Perform database migration
	database.Migrate()

	// Setup and run the Gin router
	router := api.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("couldn't run server")
	}
}
