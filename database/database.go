package database

import (
	"fmt"
	"food-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.GetString("db.user", "root"),
		config.Config.GetString("db.password", "root"),
		config.Config.GetString("db.host", "localhost"),
		config.Config.GetString("db.port", "3306"),
		config.Config.GetString("db.schema", "db"),
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)

		panic("Failed to connect to DB")
	}
}
