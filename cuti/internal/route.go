package internal

import (
	"database/sql"

	"github.com/awwabi/monorepo-academics/cuti/internal/handlers"
	"github.com/awwabi/monorepo-academics/cuti/internal/repositories"
	"github.com/awwabi/monorepo-academics/cuti/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, db *sql.DB) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(*userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(*userService)

	// prefix
	api := engine.Group("/api/cuti")

	// Routes
	api.GET("/", userHandler.SayHello)
}
