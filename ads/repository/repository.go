package repository

import (
	"context"

	"github.com/huhuhu0420/simple-ad-service/db"
	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/sirupsen/logrus"
)

type adRepository struct {
	db db.DB
}

func NewAdRepository(db db.DB) domain.AdRepository {
	return &adRepository{db}
}

func (r *adRepository) InsertNewAd(title string, startAt string, endAt string) (int, error) {
	sqlStatement := `INSERT INTO ads (title, start_at, end_at) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := r.db.QueryRow(context.Background(), sqlStatement, title, startAt, endAt).Scan(&id)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	return id, nil
}

func (r *adRepository) InsertAgeRange(id int, ageStart int, ageEnd int) error {
	sqlStatement := `INSERT INTO ad_ages (ad_id, age_start, age_end) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(context.Background(), sqlStatement, id, ageStart, ageEnd)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *adRepository) InsertCountry(id int, country []string) error {
	sqlStatement := `INSERT INTO ad_countries (ad_id, country_code) VALUES ($1, $2)`
	for _, c := range country {
		_, err := r.db.Exec(context.Background(), sqlStatement, id, c)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (r *adRepository) InsertPlatform(id int, platform []string) error {
	sqlStatement := `INSERT INTO ad_platforms (ad_id, platform) VALUES ($1, $2)`
	for _, p := range platform {
		_, err := r.db.Exec(context.Background(), sqlStatement, id, p)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (r *adRepository) InsertGender(id int, gender []string) error {
	sqlStatement := `INSERT INTO ad_genders (ad_id, gender) VALUES ($1, $2)`
	for _, g := range gender {
		_, err := r.db.Exec(context.Background(), sqlStatement, id, g)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (r *adRepository) GetAd(searchAdRequest domain.SearchAdRequest) (*domain.AdsResponse, error) {
	adsResponse := &domain.AdsResponse{}
	ad := &domain.Ad{}

	sqlStatement := `SELECT title, end_at FROM get_ads($1, $2, $3, $4, $5, $6)`
	rows, err := r.db.Query(context.Background(), sqlStatement, searchAdRequest.Offset, searchAdRequest.Limit, searchAdRequest.Age,
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
