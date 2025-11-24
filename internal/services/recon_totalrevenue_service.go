package services

import (
	"time"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/repositories"
)

type ReconTotalRevenueService struct {
	Repo *repositories.ReconTotalRevenueRepository
}

func NewReconTotalRevenueService(repo *repositories.ReconTotalRevenueRepository) *ReconTotalRevenueService {
	return &ReconTotalRevenueService{Repo: repo}
}

func (s *ReconTotalRevenueService) GetTotalRevenueByDate(date time.Time) (float64, error) {
	// repository menerima string format YYYY-MM-DD
	return s.Repo.GetTotalRevenueByDate(date.Format("2006-01-02"))
}
