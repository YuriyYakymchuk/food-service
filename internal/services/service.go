package services

import (
	"food-service/internal/models"
	"gorm.io/gorm"
)

type EntityReader interface {
	GetFoodOrders() ([]models.Food, error)
	GetFoodOrderByUser(userId string) ([]models.Food, error)
	GetUsers() ([]models.User, error)
}

type EntityWriter interface {
	CreateFoodOrder(order *models.Food) error
	UpdateFoodOrder(order *models.Food) error
	DeleteFoodOrder(id string) (models.Food, error)

	CreateUser(user *models.User) error
}

type Service interface {
	EntityReader
	EntityWriter
}

func NewService(db *gorm.DB) (Service, error) {
	return NewServiceImpl(db)
}
