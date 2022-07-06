package controllers

import (
	"encoding/json"
	"food-service/internal/database"
	"food-service/internal/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/food", getFoodOrders).Methods("GET")
	r.HandleFunc("/api/food/{userId}", getFoodOrdersByUser).Methods("GET")
	r.HandleFunc("/api/food", createFoodOrder).Methods("POST")
	r.HandleFunc("/api/food", updateFoodOrder).Methods("PUT")
	r.HandleFunc("/api/food/{id}", deleteFoodOrder).Methods("DELETE")

	r.HandleFunc("/api/user", createUser).Methods("POST")

	r.HandleFunc("/api/hello/{name}", getGreeting).Methods("GET")

	return r
}

func getFoodOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ApplicationJSON)

	var foodOrders []models.Food
	database.DB.Find(&foodOrders)
	json.NewEncoder(w).Encode(foodOrders)
}

func getFoodOrdersByUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	params := mux.Vars(r)
	var foodOrders []models.Food
	database.DB.Where(map[string]interface{}{"UserID": params["userId"]}).Find(&foodOrders)

	json.NewEncoder(w).Encode(foodOrders)
}

func createFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder models.Food
	json.NewDecoder(r.Body).Decode(&foodOrder)

	foodOrder.CreatedAt = time.Now()
	foodOrder.UpdatedAt = time.Now()
	if database.DB.Create(&foodOrder).Error != nil {
		log.Printf("Failed to create the food order: %s", foodOrder)
		message := models.Error{Message: "Failed to create the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func updateFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder models.Food
	json.NewDecoder(r.Body).Decode(&foodOrder)

	foodOrder.UpdatedAt = time.Now()
	if database.DB.Save(&foodOrder).Error != nil {
		log.Printf("Failed to update the food order: %s", foodOrder)
		message := models.Error{Message: "Failed to update the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func deleteFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r)

	var foodOrder models.Food

	if database.DB.Delete(&foodOrder, params["id"]).Error != nil {
		log.Printf("Failed to delete the food order: %s", foodOrder)
		message := models.Error{Message: "Failed to delete the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Bad request payload: %v", r.Body)
		message := models.Error{
			Message: "Bad request payload",
		}

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			log.Printf("Failed to decode the message: %v", message)
		}

		return
	}

	if database.DB.Create(&user).Error != nil {
		log.Printf("Failed to create the user: %s", user)
		message := models.Error{Message: "Failed to create the user."}
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			log.Printf("Failed to decode the message: %v", message)
		}

		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Failed to encode the user: %v", user)
	}

}

func getGreeting(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	params := mux.Vars(r)

	if len(params["name"]) < 2 {
		log.Print("Name parameter is too short.")
		message := models.Error{Message: "Name parameter should be bigger than two symbols."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode("Hello " + params["name"] + "!")
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJSON)
}
