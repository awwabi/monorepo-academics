package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/awwabi/monorepo-academics/cuti/internal/configs"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(cfg *configs.Config) *UserRepository {
	// Create PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database")

	return &UserRepository{db}
}

// Implement repository methods here
