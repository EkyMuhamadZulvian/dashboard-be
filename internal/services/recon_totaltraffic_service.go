package services

import (
	"time"

	"github.com/EkyMuhamadZulvian/dashboard-be/internal/repositories"
)

type ReconTotalTrafficService struct {
	Repo *repositories.ReconTotalTrafficRepository
}

func NewReconTotalTrafficService(repo *repositories.ReconTotalTrafficRepository) *ReconTotalTrafficService {
	return &ReconTotalTrafficService{Repo: repo}
}

func (s *ReconTotalTrafficService) GetTotalTrafficByDate(date time.Time) (int, error) {
	return s.Repo.GetTotalTrafficByDate(date.Format("2006-01-02"))
}
