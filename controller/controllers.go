package controller

import (
	"encoding/json"
	"food-service/database"
	"food-service/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

func GetFoodOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	var foodOrders []model.Food
	database.DB.Find(&foodOrders)
	json.NewEncoder(w).Encode(foodOrders)
}

func GetFoodOrdersByUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	params := mux.Vars(r)
	var foodOrders []model.Food
	database.DB.Where(map[string]interface{}{"UserID": params["userId"]}).Find(&foodOrders)

	json.NewEncoder(w).Encode(foodOrders)
}

func CreateFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder model.Food
	json.NewDecoder(r.Body).Decode(&foodOrder)

	foodOrder.CreatedAt = time.Now()
	foodOrder.UpdatedAt = time.Now()
	if database.DB.Create(&foodOrder).Error != nil {
		log.Printf("Failed to create the food order: %s", foodOrder)
		error := model.Error{Message: "Failed to create the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func UpdateFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder model.Food
	json.NewDecoder(r.Body).Decode(&foodOrder)

	foodOrder.UpdatedAt = time.Now()
	if database.DB.Save(&foodOrder).Error != nil {
		log.Printf("Failed to update the food order: %s", foodOrder)
		error := model.Error{Message: "Failed to update the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func DeleteFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r)

	var foodOrder model.Food

	if database.DB.Delete(&foodOrder, params["id"]).Error != nil {
		log.Printf("Failed to delete the food order: %s", foodOrder)
		error := model.Error{Message: "Failed to delete the food order."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(foodOrder)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if database.DB.Create(&user).Error != nil {
		log.Printf("Failed to create the user: %s", user)
		error := model.Error{Message: "Failed to create the user."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	params := mux.Vars(r)

	json.NewEncoder(w).Encode("Hello " + params["name"] + "!")
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set(contentType, applicationJSON)
}
