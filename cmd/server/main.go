package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/raul-franca/go-fc-apis/configs"
	_ "github.com/raul-franca/go-fc-apis/docs"
	"github.com/raul-franca/go-fc-apis/internal/entity"
	"github.com/raul-franca/go-fc-apis/internal/infra/database"
	"github.com/raul-franca/go-fc-apis/internal/infra/webservice/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// @title           go-fc-api
// @version         1.0
// @description     Modulo API do curso full cycle
// @description     !!!  Lembrar do Bearer no authorizations !!!
// @termsOfService  http://swagger.io/terms/

// @contact.name   Raul França
// @contact.url    https://github.com/raul-franca
// @contact.email  raulmfranca@gmail.com

// @license.name  Name License
// @license.url   https://github.com/raul-franca

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	config, err := configs.LoadConfig(".")
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
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHadler(userDB)

	//router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", config.JwtExperesIn))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
