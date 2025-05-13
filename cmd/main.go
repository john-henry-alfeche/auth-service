package main

import (
	"auth-service/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", handlers.LoginHandler())
	}

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	if err := r.Run(":3100"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
