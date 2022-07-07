package context

import (
	"errors"
	"food-service/internal/common"
	"food-service/internal/services"
	"gorm.io/gorm"
	"log"
)

var c *context

type context struct {
	beans map[string]interface{}
}

func CreateContext() {
	c = &context{beans: make(map[string]interface{})}
}

func AddBean(name string, bean interface{}) {
	c.beans[name] = bean
}

func GetService() services.Service {
	service, err := getBean(common.Service)
	if err != nil {
		return nil
	}
	return service.(services.Service)
}

func GetDB() *gorm.DB {
	db, _ := getBean(common.DB)

	return db.(*gorm.DB)
}

func getBean(name string) (interface{}, error) {
	bean, found := c.beans[name]
	if !found {
		log.Print("Bean doesn't exist: " + name)
		return nil, errors.New("Bean doesn't exist: " + name)
	}

	return bean, nil
}
