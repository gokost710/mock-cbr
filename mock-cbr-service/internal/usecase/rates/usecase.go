package ratesusecase

import (
	"fmt"
	"net/http"
	"time"
)

type RatesRepository interface {
	GetDailyXML(date time.Time) ([]byte, error)
}

type Usecase struct {
	repo RatesRepository
}

func New(repo RatesRepository) *Usecase {
	return &Usecase{repo: repo}
}

// GetDailyXML возвращает XML курсов за дату из параметра date_req (формат dd/mm/yyyy)
func (u *Usecase) GetDailyXML(dateReq string) ([]byte, int, error) {
	if dateReq == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("date_req is required")
	}
	t, err := time.Parse("02/01/2006", dateReq)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid date_req format, want dd/mm/yyyy")
	}
	data, err := u.repo.GetDailyXML(t)
	if err != nil {
		return nil, http.StatusBadGateway, err
	}
	return data, http.StatusOK, nil
}
