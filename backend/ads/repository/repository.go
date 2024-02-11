package repository

import (
	"database/sql"

	"github.com/huhuhu0420/simple-ad-service/domain"
)

type adRepository struct {
	db *sql.DB
}

func NewAdRepository(db *sql.DB) domain.AdRepository {
	return &adRepository{db}
}

func (r *adRepository) CreateAd(ad domain.AdInfo, conditions domain.Conditions) error {
	return nil
}

func (r *adRepository) GetAd(searchAdRequest domain.SearchAdRequest) (domain.AdsResponse, error) {
	return domain.AdsResponse{}, nil
}
