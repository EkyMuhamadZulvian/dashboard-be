package routes

import (
	"database/sql"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, db *sql.DB) {
	r.Get("/api/hello", handlers.HelloHandler)
	r.Get("/api/recon/totalrevenue", handlers.ReconTotalRevenueHandler(db))
	r.Get("/api/recon/totaltraffic", handlers.ReconTotalTrafficHandler(db))
}
