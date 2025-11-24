package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/repositories"
	"github.com/EkyMuhamadZulvian/dashboard-be/internal/services"
)

func ReconTotalRevenueHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ambil query param ?date=YYYY-MM-DD
		dateStr := r.URL.Query().Get("date")
		if dateStr == "" {
			dateStr = time.Now().Format("2006-01-02")
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		repo := repositories.NewReconTotalRevenueRepository(db)
		service := services.NewReconTotalRevenueService(repo)

		totalRevenue, err := service.GetTotalRevenueByDate(date)
		if err != nil {
			http.Error(w, "Failed to get total revenue", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"total_revenue": totalRevenue})
	}
}
