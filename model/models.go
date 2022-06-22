package model

import (
	"time"
)

type Food struct {
	ID        int       `json:"id" gorm:"column:ID"`
	UserID    int       `json:"userId" gorm:"column:UserID"`
	Name      string    `json:"name" gorm:"size:255;column:Name"`
	Price     float64   `json:"price" gorm:"column:Price"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UpdatedAt"`
}

type User struct {
	ID        int       `json:"id,primary_key" gorm:"column:ID"`
	FirstName string    `json:"firstName" gorm:"size:255;column:FirstName"`
	LastName  string    `json:"lastName" gorm:"size:255;column:LastName"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UpdatedAt"`
}

type Error struct {
	Message string `json:"message"`
}

func (Food) TableName() string {
	return "Food"
}

func (User) TableName() string {
	return "User"
}
