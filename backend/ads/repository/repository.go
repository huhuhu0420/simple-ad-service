package repository

import (
	"database/sql"

	"github.com/huhuhu0420/ads-service/domain"
)

type adRepository struct {
	db *sql.DB
}

func NewAdRepository(db *sql.DB) domain.AdRepository {
	return &adRepository{db}
}

func (r *adRepository) CreateAd(ad domain.AdInfo) error {
	return nil
}

func (r *adRepository) GetAd(conditions domain.Conditions) (domain.AdsResponse, error) {
	return domain.AdsResponse{}, nil
}
