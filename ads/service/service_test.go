package service

import (
	"testing"

	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/huhuhu0420/simple-ad-service/domain/mocks"
	"github.com/stretchr/testify/suite"
)

type serviceSuite struct {
	suite.Suite
	service    domain.AdService
	repository *mocks.AdRepository
}

func (s *serviceSuite) SetupSuite() {
	s.repository = new(mocks.AdRepository)
	s.service = NewAdService(s.repository)
}

func (s *serviceSuite) TestCreateAd() {
	s.repository.On("InsertNewAd", "test", "2021-01-01", "2021-01-02").Return(1, nil)
	s.repository.On("InsertAgeRange", 1, 10, 20).Return(nil)
	s.repository.On("InsertCountry", 1, []string{"US"}).Return(nil)
	s.repository.On("InsertPlatform", 1, []string{"ios"}).Return(nil)
	s.repository.On("InsertGender", 1, []string{"M"}).Return(nil)

	err := s.service.CreateAd(domain.AdInfo{
		Title:   "test",
		StartAt: "2021-01-01",
		EndAt:   "2021-01-02",
	}, domain.Conditions{
		AgeStart: 10,
		AgeEnd:   20,
		Country:  []string{"US"},
		Platform: []string{"ios"},
		Gender:   []string{"M"},
	})
	s.NoError(err)
}

func (s *serviceSuite) TestGetAd() {
	// test if the service calls the repository
	s.repository.On("GetAd", domain.SearchAdRequest{}).Return(&domain.AdsResponse{}, nil)
	_, err := s.service.GetAd(domain.SearchAdRequest{})
	s.NoError(err)
}

func TestService(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}
