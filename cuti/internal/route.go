package internal

import (
	"github.com/awwabi/monorepo-academics/cuti/internal/configs"
	"github.com/awwabi/monorepo-academics/cuti/internal/handlers"
	"github.com/awwabi/monorepo-academics/cuti/internal/repositories"
	"github.com/awwabi/monorepo-academics/cuti/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) *gin.Engine {
	// Load configurations
	cfg := configs.LoadConfig()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(cfg)

	// Initialize services
	userService := services.NewUserService(*userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(*userService)

	// prefix
	api := engine.Group("/api/cuti")

	// Routes
	api.GET("/", userHandler.SayHello)

	return engine
}
