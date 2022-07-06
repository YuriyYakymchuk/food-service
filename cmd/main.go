package main

import (
	"fmt"
	"food-service/internal/config"
	"food-service/internal/controllers"
	"food-service/internal/database"
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

	config.LoadProperties("resources/food.properties")
	database.InitDB()
	r := controllers.InitializeRoutes()
	startServer(r)
}

func startServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",
		config.Config.GetString("host", "localhost"), config.Config.GetString("port", "8080")),
		handlers.LoggingHandler(os.Stdout, router)))
}
