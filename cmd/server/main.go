package main

import (
	"fmt"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/raul-franca/go-fc-apis/configs"
	"github.com/raul-franca/go-fc-apis/internal/entity"
	"github.com/raul-franca/go-fc-apis/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("failed to load config file, error: %v", err))
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database, error: %v", err))
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate the schema, error: %v", err))
	}
	productDB := database.NewProduct(db)

}
