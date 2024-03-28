package repository

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/huhuhu0420/simple-ad-service/db"
	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	suite.Suite
	mock       pgxmock.PgxPoolIface
	mockCache  redismock.ClientMock
	repository domain.AdRepository
}

func MockDB() (db.DB, error) {
	mockPool, err := pgxmock.NewPool()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return mockPool, nil
}

func (s *repositorySuite) SetupTest() {
	var err error
	var db db.DB
	s.mock, err = pgxmock.NewPool()
	cache, b := redismock.NewClientMock()
	s.mockCache = b
	db = s.mock
	if err != nil {
		s.T().Fatal(err)
	}
	s.repository = NewAdRepository(db, cache)
}

func (s *repositorySuite) TearDownTest() {
	s.repository.(*adRepository).db.Close()
}

func (s *repositorySuite) TestInsertNewAd() {
	s.mock.ExpectQuery("INSERT INTO ads").
		WithArgs("test", "2021-01-01", "2021-01-02").
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

	id, err := s.repository.InsertNewAd("test", "2021-01-01", "2021-01-02")
	s.NoError(err)
	s.Assert().Equal(1, id)
}

func (s *repositorySuite) TestInsertAgeRange() {
	s.mock.ExpectExec("INSERT INTO ad_ages").
		WithArgs(1, 10, 20).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.repository.InsertAgeRange(1, 10, 20)
	s.NoError(err)
}

func (s *repositorySuite) TestInsertCountry() {
	s.mock.ExpectExec("INSERT INTO ad_countries").
		WithArgs(1, "US").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.repository.InsertCountry(1, []string{"US"})
	s.NoError(err)
}

func (s *repositorySuite) TestInsertPlatform() {
	s.mock.ExpectExec("INSERT INTO ad_platforms").
		WithArgs(1, "ios").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.repository.InsertPlatform(1, []string{"ios"})
	s.NoError(err)
}

func (s *repositorySuite) TestInsertGender() {
	s.mock.ExpectExec("INSERT INTO ad_genders").
		WithArgs(1, "M").
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	err := s.repository.InsertGender(1, []string{"M"})
	s.NoError(err)
}

func (s *repositorySuite) TestGetAd() {
	request := domain.SearchAdRequest{
		Offset:   0,
		Limit:    10,
		Age:      18,
		Country:  "US",
		Platform: "ios",
		Gender:   "M",
	}

	rows := pgxmock.NewRows([]string{"title", "end_at"}).
		AddRow("test", "2021-01-01").
		AddRow("test2", "2021-01-02")
	s.mock.ExpectQuery("SELECT").
		WithArgs(0, 10, 18, "M", "US", "ios").
		WillReturnRows(rows)

	key, _ := json.Marshal(request)
	s.mockCache.ExpectGet(string(key)).RedisNil()
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	ttl := midnight.Sub(now)
	s.mockCache.Regexp().ExpectSet(string(key), `.*`, ttl).SetVal("OK")
	ads, err := s.repository.GetAd(request)
	s.NoError(err)
	s.Assert().Equal(2, len(ads.Items))
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(repositorySuite))
}
