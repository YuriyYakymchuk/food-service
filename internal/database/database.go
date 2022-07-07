package database

import (
	"food-service/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(config.ConstructDataBaseConnectionString()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalln("Failed to initialize the DB")
		return nil, err
	}

	return db, nil
}
