package main

import (
	"log"
	"net/http"

	"github.com/devfullcycle/goexpert/9-API/configs"
	_ "github.com/devfullcycle/goexpert/9-API/docs"
	"github.com/devfullcycle/goexpert/9-API/internal/entity"
	"github.com/devfullcycle/goexpert/9-API/internal/infra/database"
	"github.com/devfullcycle/goexpert/9-API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Vinicius Souza
// @contact.url    http://www.intellisys.com.br
// @contact.email  sac@intellisys.com.br

// @license.name  Intellisys Informática
// @license.url   https://www.intellisys.com.br

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Use(middleware.Recoverer)                                       // Determina o Reinício do Sistema em caso de Encerramento por Falha
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))             // Captura Dados para Validação do Token
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JWTExpiresIn)) // Captura Dados para Validação da Expiração do Token

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
