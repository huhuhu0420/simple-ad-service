package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/huhuhu0420/simple-ad-service/domain/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	handler *AdHandler
	service *mocks.AdService
	router  *gin.Engine
}

func (s *handlerSuite) SetupTest() {
	s.service = new(mocks.AdService)
	s.handler = &AdHandler{
		Service: s.service,
	}
	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	NewAdHandler(s.router, s.service)
}

func (s *handlerSuite) TestGetAdResponse200() {
	s.service.On("GetAd", mock.Anything).Return(&domain.AdsResponse{}, nil)
	req, _ := http.NewRequest("GET", "/api/v1/ad", nil)
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
}

func (s *handlerSuite) TestGetAdResponse500() {
	s.service.On("GetAd", mock.Anything).Return(nil, errors.New("error"))
	req, _ := http.NewRequest("GET", "/api/v1/ad", nil)
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
}

func (s *handlerSuite) TestGetAdResponse400() {
	req, _ := http.NewRequest("GET", "/api/v1/ad?Gender=2", nil)
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *handlerSuite) TestCreateAdResponse200() {
	s.service.On("CreateAd", mock.Anything, mock.Anything).Return(nil)

	adInfo := domain.AdInfo{
		Title:   "test",
		StartAt: "2023-01-01T00:00:00Z",
		EndAt:   "2023-12-01T00:00:00Z",
	}
	data, _ := json.Marshal(adInfo)
	req, _ := http.NewRequest("POST", "/api/v1/ad", bytes.NewReader(data))

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
}

func (s *handlerSuite) TestCreateAdResponse400() {
	adInfo := domain.AdInfo{
		Title:   "test",
		StartAt: "2023-01-01T00:00:00Z",
		EndAt:   "2023-12-01T00:00:00Z",
		Conditions: domain.Conditions{
			AgeStart: 18,
			AgeEnd:   65,
			Gender:   []string{"L"},
		},
	}
	data, _ := json.Marshal(adInfo)
	req, _ := http.NewRequest("POST", "/api/v1/ad", bytes.NewReader(data))

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *handlerSuite) TestCreateAdResponse500() {
	s.service.On("CreateAd", mock.Anything, mock.Anything).Return(errors.New("error"))

	adInfo := domain.AdInfo{
		Title:   "test",
		StartAt: "2023-01-01T00:00:00Z",
		EndAt:   "2023-12-01T00:00:00Z",
	}
	data, _ := json.Marshal(adInfo)
	req, _ := http.NewRequest("POST", "/api/v1/ad", bytes.NewReader(data))

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}
