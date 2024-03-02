package handler

import (
	"testing"

	"github.com/huhuhu0420/simple-ad-service/domain/mocks"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	handler *AdHandler
	service *mocks.AdService
}

func (s *handlerSuite) SetupSuite() {
	s.service = new(mocks.AdService)
	s.handler = &AdHandler{
		Service: s.service,
	}
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}
