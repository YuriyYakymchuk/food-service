package database

import (
	"food-service/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(config.ConstructDataBaseConnectionString()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)

		panic("Failed to connect to DB")
	}
}
