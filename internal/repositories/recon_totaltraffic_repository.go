package repositories

import (
	"database/sql"
	"log"
)

type ReconTotalTrafficRepository struct {
	DB *sql.DB
}

func NewReconTotalTrafficRepository(db *sql.DB) *ReconTotalTrafficRepository {
	return &ReconTotalTrafficRepository{DB: db}
}

// GetTotalTrafficByDate mengembalikan total Traffic untuk tanggal tertentu
func (r *ReconTotalTrafficRepository) GetTotalTrafficByDate(date string) (int, error) {
	var total int

	query := `
        SELECT SUM(a.Traffic)
        FROM Analytic.Summary a
        WHERE a.Traffic IS NOT NULL AND a.Date = ?
    `

	err := r.DB.QueryRow(query, date).Scan(&total)
	if err != nil {
		log.Println("Error GetTotalTrafficByDate:", err)
		return 0, err
	}

	return total, nil
}
