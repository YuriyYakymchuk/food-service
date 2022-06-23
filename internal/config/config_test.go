package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructDataBaseConnectionString(t *testing.T) {
	LoadProperties("food_test.properties")

	expected := "user:pwd@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	actual := ConstructDataBaseConnectionString()

	assert.Equal(t, expected, actual, fmt.Sprintf("Wrong DB connection string: %s", actual))
}

func TestConstructDataBaseConnectionStringDefault(t *testing.T) {
	LoadProperties("food_test_default.properties")

	expected := ":@tcp(:)/?charset=utf8mb4&parseTime=True&loc=Local"
	actual := ConstructDataBaseConnectionString()

	assert.Equal(t, expected, actual, fmt.Sprintf("Wrong DB connection string: %s", actual))
}
