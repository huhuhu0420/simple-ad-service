package service

import "github.com/huhuhu0420/simple-ad-service/domain"

type adService struct {
	repo domain.AdRepository
}

func NewAdService(repo domain.AdRepository) domain.AdService {
	return &adService{repo}
}

func (s *adService) CreateAd(ad domain.AdInfo, conditions domain.Conditions) error {
	return s.repo.CreateAd(ad, conditions)
}

func (s *adService) GetAd(searchAdRequest domain.SearchAdRequest) (*domain.AdsResponse, error) {
	return s.repo.GetAd(searchAdRequest)
}
