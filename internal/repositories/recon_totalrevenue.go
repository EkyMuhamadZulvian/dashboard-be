package repositories

import (
	"database/sql"
	"log"
)

type ReconTotalRevenueRepository struct {
	DB *sql.DB
}

func NewReconTotalRevenueRepository(db *sql.DB) *ReconTotalRevenueRepository {
	return &ReconTotalRevenueRepository{DB: db}
}

// GetTotalRevenueByDate mengembalikan total revenue untuk tanggal tertentu
func (r *ReconTotalRevenueRepository) GetTotalRevenueByDate(date string) (float64, error) {
	var total float64

	query := `
        SELECT SUM(a.Amount) 
        FROM Analytic.Summary a
        WHERE a.Amount IS NOT NULL AND a.Date = ?
    `

	err := r.DB.QueryRow(query, date).Scan(&total)
	if err != nil {
		log.Println("Error GetTotalRevenueByDate:", err)
		return 0, err
	}

	return total, nil
}
