package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Connect to the database
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set up database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established.")
}

// MigrateDatabase migrates the database schema
func MigrateDatabase() {
	err := DB.AutoMigrate(&Item{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}

// SeedDatabase populates the database with initial data
func SeedDatabase() {
	// Define initial seed data
	items := []Item{
		{
			Name:   "Potato",
			Price:  10,
			Amount: 200,
		},
		{
			Name:   "Carrot",
			Price:  15,
			Amount: 50,
		},
	}

	if DB.Where("name = ?", "Potato").First(&Item{}).Error == nil {
		return
	}

	// Insert seed data into the database
	for _, item := range items {
		err := DB.Create(&item).Error
		if err != nil {
			log.Printf("Failed to seed notification: %v", err)
		} else {
			log.Printf("Seeded notification: %+v", item)
		}
	}
	log.Println("Database seeding completed.")
}

type Item struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Amount    int       `json:"amount"`
}

func main() {
	// Initialize database and run migrations
	InitDB()
	MigrateDatabase()
	SeedDatabase()

	r := mux.NewRouter()
	// Item Handlers
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		var items []Item
		err := DB.Find(&items).Error

		if err != nil {
			http.Error(w, "Failed to read items", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(items)
	}).Methods("GET")

	// Missing: Implement create, update and delete

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
