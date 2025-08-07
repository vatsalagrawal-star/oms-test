package main

import (
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
	router.Run(":8080")
}
