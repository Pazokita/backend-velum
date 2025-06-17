package main

import (
	"database/sql"
	"log"
	"velum/internal/maps"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database/maps.db")
	if err != nil {
		log.Fatal("Échec connexion DB:", err)
	}
	defer db.Close()

	repo := maps.NewRepository(db)
	service := maps.NewService(repo)
	handler := maps.NewHandler(service)

	r := gin.Default()

	r.GET("/maps", handler.GetMaps)
	r.GET("/maps/:id", handler.GetMapByID)

	log.Println("Serveur lancé sur http://localhost:8080")
	r.Run(":8080")
}
