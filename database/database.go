package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func InitDB() {

	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)

		panic("Failed to connect to DB")
	}
}
