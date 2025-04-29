package main

import (
	"velum/internal/maps"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route GET /maps
	r.GET("/maps", maps.GetMapsHandler)

	// Lancement du serveur sur le port 8080
	r.Run(":8080")
}
