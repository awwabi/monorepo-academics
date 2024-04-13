package main

import (
	"github.com/awwabi/monorepo-academics/cuti/internal"
	"github.com/awwabi/monorepo-academics/cuti/internal/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()
	db := configs.InitDB()

	// Initialize Gin router
	r := gin.Default()

	internal.RegisterRoutes(r, db)

	// Run the server
	r.Run(":8080")
}
