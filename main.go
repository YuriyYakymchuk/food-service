package main

import (
	"food-service/database"
	"food-service/route"
)

func main() {

	database.InitDB()
	route.InitializeRoutes()

}
