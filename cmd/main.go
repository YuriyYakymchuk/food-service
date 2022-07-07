package main

import (
	"fmt"
	"food-service/internal/common"
	"food-service/internal/config"
	context2 "food-service/internal/context"
	"food-service/internal/controllers"
	"food-service/internal/database"
	"food-service/internal/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Book struct {
	Name  string
	Price float64
}

type PrintFunction func(text string)

func (f PrintFunction) printName(book Book) {
	f(book.Name)
}

func printI(name string) {
	fmt.Println("Name inside:", name)
}

func printO(name string) {
	fmt.Println("Name outside:", name)
}

type A interface {
	a() string
}

type B interface {
	b() int
}

type C interface {
	A
	B
}

type Impl struct {
}

func (impl *Impl) a() string {
	return "a"
}

func (impl *Impl) b() int64 {
	return 1
}

func main() {
	log.Print("Starting the food order service ...")

	status, err := run()
	if err != nil {
		log.Print("Failed to start the service.")
		os.Exit(status)
	}
}

func run() (int, error) {
	config.LoadProperties("resources/food.properties")

	context2.CreateContext()

	db, err := database.InitDB()
	if err != nil {
		return 1, err
	}
	context2.AddBean(common.DB, db)

	service, err := services.NewService(db)
	if err != nil {
		return 1, err
	}
	context2.AddBean(common.Service, service)

	r := controllers.InitializeRoutes()
	startServer(r)

	return 0, nil
}

func startServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",
		config.Config.GetString("host", "localhost"), config.Config.GetString("port", "8080")),
		handlers.LoggingHandler(os.Stdout, router)))
}
