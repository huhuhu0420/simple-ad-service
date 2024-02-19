package service

import (
	"fmt"

	"github.com/huhuhu0420/simple-ad-service/domain"
)

type adService struct {
	repo domain.AdRepository
}

func NewAdService(repo domain.AdRepository) domain.AdService {
	return &adService{repo}
}

func (s *adService) CreateAd(ad domain.AdInfo, conditions domain.Conditions) error {
	id, err := s.repo.InsertNewAd(ad.Title, ad.StartAt, ad.EndAt)
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func (s *adService) GetAd(searchAdRequest domain.SearchAdRequest) (*domain.AdsResponse, error) {
	return s.repo.GetAd(searchAdRequest)
}
