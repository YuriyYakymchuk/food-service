package config

import (
	"fmt"
	"github.com/magiconair/properties"
)

var Config *properties.Properties

func LoadProperties(file string) {
	Config = properties.MustLoadFile(file, properties.UTF8)
}

func ConstructDataBaseConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.GetString("db.user", ""),
		Config.GetString("db.password", ""),
		Config.GetString("db.host", ""),
		Config.GetString("db.port", ""),
		Config.GetString("db.schema", ""),
	)
}
