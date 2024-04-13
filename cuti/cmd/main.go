package main

import (
	"github.com/awwabi/monorepo-academics/cuti/internal"
	"github.com/awwabi/monorepo-academics/cuti/internal/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()

	// Initialize Gin router
	r := gin.Default()

	internal.RegisterRoutes(r)

	// Run the server
	r.Run(":8080")
}
