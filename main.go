package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"secret-santa-backend/models"
	"secret-santa-backend/routes"
)

var db *gorm.DB
var googleOauthConfig = &oauth2.Config{}

func main() {

	initDB()

	// Auto migrate models
	db.AutoMigrate(&models.User{}, &models.Group{}, &models.Assignment{})

	// Initialize router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router, db)

	// Start server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
func initGoogleAuth() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"}, // Adjust scopes as needed
		Endpoint:     google.Endpoint,
	}
}
