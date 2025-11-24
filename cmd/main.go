package main

import (
	"log"
	"net/http"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/config"
	"github.com/EkyMuhamadZulvian/dashboard-be/internal/database"
	"github.com/EkyMuhamadZulvian/dashboard-be/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	r := chi.NewRouter()
	routes.RegisterRoutes(r, db)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
