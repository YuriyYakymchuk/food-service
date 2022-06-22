package main

import (
	"food-service/config"
	"food-service/database"
	"food-service/route"
)

func main() {

	config.LoadProperties()
	database.InitDB()
	route.InitializeRoutes()

}
