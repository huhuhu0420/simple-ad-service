package service

import (
	"github.com/huhuhu0420/simple-ad-service/domain"
)

type adService struct {
	repo domain.AdRepository
}

func NewAdService(repo domain.AdRepository) domain.AdService {
	return &adService{repo}
}

func (s *adService) insertConditions(id int, conditions domain.Conditions) error {
	if conditions.AgeStart != 0 && conditions.AgeEnd != 0 {
		if err := s.repo.InsertAgeRange(id, conditions.AgeStart, conditions.AgeEnd); err != nil {
			return err
		}
	}
	if len(conditions.Country) > 0 {
		if err := s.repo.InsertCountry(id, conditions.Country); err != nil {
			return err
		}
	}
	if len(conditions.Platform) > 0 {
		if err := s.repo.InsertPlatform(id, conditions.Platform); err != nil {
			return err
		}
	}
	if len(conditions.Gender) > 0 {
		if err := s.repo.InsertGender(id, conditions.Gender); err != nil {
			return err
		}
	}
	return nil
}

func (s *adService) CreateAd(ad domain.AdInfo, conditions domain.Conditions) error {
	id, err := s.repo.InsertNewAd(ad.Title, ad.StartAt, ad.EndAt)
	if err != nil {
		return err
	}
	if err := s.insertConditions(id, conditions); err != nil {
		return err
	}
	return nil
}

func (s *adService) GetAd(searchAdRequest domain.SearchAdRequest) (*domain.AdsResponse, error) {
	ad, err := s.repo.GetAd(searchAdRequest)
	if err != nil {
		return nil, err
	}
	return ad, nil
}
