package config

import "github.com/magiconair/properties"

var Config *properties.Properties

func LoadProperties() {
	Config = properties.MustLoadFile("resources/food.properties", properties.UTF8)
}
