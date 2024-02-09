package service

import "github.com/huhuhu0420/ads-service/domain"

type adService struct {
	repo domain.AdRepository
}

func NewAdService(repo domain.AdRepository) domain.AdService {
	return &adService{repo}
}

func (s *adService) CreateAd(ad domain.AdInfo) error {
	return s.repo.CreateAd(ad)
}

func (s *adService) GetAd(conditions domain.Conditions) (domain.AdsResponse, error) {
	return s.repo.GetAd(conditions)
}
