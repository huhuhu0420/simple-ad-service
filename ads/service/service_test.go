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
	s.repository.On("InsertCountry", 1, []string{"US", "TW", "JP"}).Return(nil)
	s.repository.On("InsertPlatform", 1, []string{"ios", "web"}).Return(nil)
	s.repository.On("InsertGender", 1, []string{"M"}).Return(nil)
	s.repository.On("InvalidateAllCache").Return(nil)

	testCase := []domain.AdInfo{}
	// ad has conditions
	testCase = append(testCase, domain.AdInfo{
		Title:   "test",
		StartAt: "2021-01-01",
		EndAt:   "2021-01-02",
		Conditions: domain.Conditions{
			AgeStart: 10,
			AgeEnd:   20,
			Country:  []string{"US", "TW", "JP"},
			Platform: []string{"ios", "web"},
			Gender:   []string{"M"},
		},
	})
	// ad has no conditions
	testCase = append(testCase, domain.AdInfo{
		Title:   "test",
		StartAt: "2021-01-01",
		EndAt:   "2021-01-02",
		Conditions: domain.Conditions{
			AgeStart: 0,
			AgeEnd:   0,
			Country:  []string{},
			Platform: []string{},
			Gender:   []string{},
		},
	})

	for _, tc := range testCase {
		err := s.service.CreateAd(tc, tc.Conditions)
		s.NoError(err)
	}
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
