package services

import (
	"food-service/internal/models"
	"gorm.io/gorm"
)

type serviceImpl struct {
	db *gorm.DB
}

func NewServiceImpl(db *gorm.DB) (Service, error) {
	return &serviceImpl{db: db}, nil
}

func (s *serviceImpl) GetFoodOrders() ([]models.Food, error) {
	var orders []models.Food
	err := s.db.Find(&orders).Error

	return orders, err
}

func (s *serviceImpl) GetFoodOrderByUser(userId string) ([]models.Food, error) {
	var orders []models.Food
	err := s.db.Where(map[string]string{"UserID": userId}).Find(&orders).Error

	return orders, err
}

func (s *serviceImpl) CreateFoodOrder(order *models.Food) error {
	return s.db.Create(order).Error
}

func (s *serviceImpl) UpdateFoodOrder(order *models.Food) error {
	var foodOrder models.Food
	err := s.db.First(&foodOrder, order.ID).Error
	if err != nil {
		return err
	}

	order.CreatedAt = foodOrder.CreatedAt

	return s.db.Updates(order).Error
}

func (s *serviceImpl) DeleteFoodOrder(id string) (models.Food, error) {
	var order models.Food
	err := s.db.First(&order, id).Error
	if err != nil {
		return order, err
	}
	err = s.db.Delete(&order, id).Error

	return order, err
}

func (s *serviceImpl) GetUsers() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error

	return users, err
}

func (s *serviceImpl) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}
