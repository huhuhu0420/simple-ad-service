package repository

import (
	"database/sql"

	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/sirupsen/logrus"
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

func (r *adRepository) GetAd(searchAdRequest domain.SearchAdRequest) (*domain.AdsResponse, error) {
	adsResponse := &domain.AdsResponse{}
	ad := &domain.Ad{}

	sqlStatement := `SELECT title, end_at FROM get_ads($1, $2, $3, $4, $5, $6)`
	rows, err := r.db.Query(sqlStatement, searchAdRequest.Offset, searchAdRequest.Limit, searchAdRequest.Age,
		searchAdRequest.Gender, searchAdRequest.Country, searchAdRequest.Platform)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&ad.Title, &ad.EndAt)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		adsResponse.Items = append(adsResponse.Items, *ad)
	}

	return adsResponse, nil
}
