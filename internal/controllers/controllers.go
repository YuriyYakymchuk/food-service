package controllers

import (
	"encoding/json"
	context2 "food-service/internal/context"
	"food-service/internal/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"

	FailedToEncodeResponse = "Failed to encode response."
	InternalServerError    = "Internal server error."
	FailedToRetrieveOrders = "Failed to retrieve orders."
	BadRequestPayload      = "Bad request payload."
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/food", getFoodOrders).Methods("GET")
	r.HandleFunc("/api/food/{userId}", getFoodOrdersByUser).Methods("GET")
	r.HandleFunc("/api/food", createFoodOrder).Methods("POST")
	r.HandleFunc("/api/food", updateFoodOrder).Methods("PUT")
	r.HandleFunc("/api/food/{id}", deleteFoodOrder).Methods("DELETE")

	r.HandleFunc("/api/user", getUsers).Methods("GET")
	r.HandleFunc("/api/user", createUser).Methods("POST")

	r.HandleFunc("/api/hello/{name}", getGreeting).Methods("GET")

	return r
}

func getFoodOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ApplicationJSON)

	service := context2.GetService()
	if service == nil {
		handleInternalServerError(w, InternalServerError)
		return
	}

	orders, err := service.GetFoodOrders()
	if err != nil {
		handleInternalServerError(w, FailedToRetrieveOrders)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		handleInternalServerError(w, FailedToEncodeResponse)
		return
	}
}

func getFoodOrdersByUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	params := mux.Vars(r)
	var foodOrders []models.Food

	service := context2.GetService()
	if service == nil {
		handleInternalServerError(w, InternalServerError)
		return
	}

	foodOrders, err := service.GetFoodOrderByUser(params["userId"])
	if err != nil {
		handleInternalServerError(w, FailedToRetrieveOrders)
		return
	}

	err = json.NewEncoder(w).Encode(foodOrders)
	if err != nil {
		handleInternalServerError(w, FailedToEncodeResponse)
		return
	}
}

func createFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder models.Food
	err := json.NewDecoder(r.Body).Decode(&foodOrder)

	if err != nil {
		handleBadRequest(w, BadRequestPayload)
		return
	}

	service := context2.GetService()
	if service == nil {
		handleInternalServerError(w, InternalServerError)
		return
	}

	err = service.CreateFoodOrder(&foodOrder)

	if err != nil {
		log.Printf("Failed to create the food order: %s", foodOrder)
		handleInternalServerError(w, "Failed to create the food order.")
		return
	}

	err = json.NewEncoder(w).Encode(foodOrder)
	if err != nil {
		handleInternalServerError(w, FailedToEncodeResponse)
		return
	}
}

func updateFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var foodOrder models.Food
	err := json.NewDecoder(r.Body).Decode(&foodOrder)

	if err != nil {
		handleBadRequest(w, BadRequestPayload)
	}

	service := context2.GetService()
	if service == nil {
		handleInternalServerError(w, InternalServerError)
		return
	}

	err = service.UpdateFoodOrder(&foodOrder)

	if err != nil {
		log.Printf("Failed to update the food order: %s", foodOrder)
		handleRecordNotFound(w)
		return
	}

	err = json.NewEncoder(w).Encode(foodOrder)
	if err != nil {
		handleInternalServerError(w, FailedToEncodeResponse)
		return
	}
}

func deleteFoodOrder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r)

	service := context2.GetService()
	if service == nil {
		handleInternalServerError(w, InternalServerError)
		return
	}

	foodOrder, err := service.DeleteFoodOrder(params["id"])
	if err != nil {
		log.Printf("Failed tp delete the food order with ID: %s", params["id"])
		handleRecordNotFound(w)
		return
	}

	err = json.NewEncoder(w).Encode(foodOrder)
	if err != nil {
		handleInternalServerError(w, FailedToEncodeResponse)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	service := context2.GetService()
	if service == nil {
		message := models.Error{Message: "Internal server error."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	users, err := service.GetUsers()
	if err != nil {
		message := models.Error{Message: "Failed to retrieve users."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		message := models.Error{Message: "Failed to encode response."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleBadRequest(w, BadRequestPayload)
		return
	}

	service := context2.GetService()
	if service == nil {
		message := models.Error{Message: "Internal server error."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}

	err = service.CreateUser(&user)

	if err != nil {
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

func handleBadRequest(w http.ResponseWriter, text string) {
	log.Print(text)
	message := models.Error{
		Message: text,
	}

	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("Failed to decode the message: %v", message)
	}
}

func handleInternalServerError(w http.ResponseWriter, message string) {
	response := models.Error{Message: message}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

func handleRecordNotFound(w http.ResponseWriter) {
	response := models.Error{Message: "Record not found"}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
