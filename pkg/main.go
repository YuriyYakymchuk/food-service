package main

import (
	"fmt"
	"food-service/pkg/config"
	"food-service/pkg/controller"
	"food-service/pkg/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	config.LoadProperties("resources/food.properties")
	database.InitDB()
	r := controller.InitializeRoutes()
	startServer(r)
}

func startServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",
		config.Config.GetString("host", "localhost"), config.Config.GetString("port", "8080")),
		handlers.LoggingHandler(os.Stdout, router)))
}
