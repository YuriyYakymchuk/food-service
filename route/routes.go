package route

import (
	"fmt"
	"food-service/config"
	"food-service/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitializeRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/api/food", controller.GetFoodOrders).Methods("GET")
	r.HandleFunc("/api/food/{userId}", controller.GetFoodOrdersByUser).Methods("GET")
	r.HandleFunc("/api/food", controller.CreateFoodOrder).Methods("POST")
	r.HandleFunc("/api/food", controller.UpdateFoodOrder).Methods("PUT")
	r.HandleFunc("/api/food/{id}", controller.DeleteFoodOrder).Methods("DELETE")

	r.HandleFunc("/api/user", controller.CreateUser).Methods("POST")

	r.HandleFunc("/api/hello/{name}", controller.GetGreeting).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",
		config.Config.GetString("host", "localhost"), config.Config.GetString("port", "8080")), r))
}
