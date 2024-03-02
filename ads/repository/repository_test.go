package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huhuhu0420/simple-ad-service/domain"
	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository domain.AdRepository
}

func (s *repositorySuite) SetupTest() {
	var db *sql.DB
	var err error
	db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatal(err)
	}
	s.repository = NewAdRepository(db)
}

func (s *repositorySuite) TearDownTest() {
	s.repository.(*adRepository).db.Close()
}

func (s *repositorySuite) TestInsertNewAd() {
	s.mock.ExpectQuery("INSERT INTO ads").
		WithArgs("test", "2021-01-01", "2021-01-02").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := s.repository.InsertNewAd("test", "2021-01-01", "2021-01-02")
	s.NoError(err)
	s.Assert().Equal(1, id)
}

func (s *repositorySuite) TestInsertAgeRange() {
	s.mock.ExpectExec("INSERT INTO ad_ages").
		WithArgs(1, 10, 20).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.InsertAgeRange(1, 10, 20)
	s.NoError(err)
}

func (s *repositorySuite) TestInsertCountry() {
	s.mock.ExpectExec("INSERT INTO ad_countries").
		WithArgs(1, "US").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.InsertCountry(1, []string{"US"})
	s.NoError(err)
}

func (s *repositorySuite) TestInsertPlatform() {
	s.mock.ExpectExec("INSERT INTO ad_platforms").
		WithArgs(1, "ios").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.InsertPlatform(1, []string{"ios"})
	s.NoError(err)
}

func (s *repositorySuite) TestInsertGender() {
	s.mock.ExpectExec("INSERT INTO ad_genders").
		WithArgs(1, "M").
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := s.repository.InsertGender(1, []string{"M"})
	s.NoError(err)
}

func (s *repositorySuite) TestGetAd() {
	rows := sqlmock.NewRows([]string{"title", "end_at"}).
		AddRow("test", "2021-01-01").
		AddRow("test2", "2021-01-02")
	s.mock.ExpectQuery("SELECT").
		WillReturnRows(rows)

	ads, err := s.repository.GetAd(domain.SearchAdRequest{})
	s.NoError(err)
	s.Assert().Equal(2, len(ads.Items))
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(repositorySuite))
}
