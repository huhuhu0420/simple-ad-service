package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInsertNewAd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("INSERT INTO ads").
		WithArgs("test", "2021-01-01", "2021-01-02").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	r := NewAdRepository(db)
	id, err := r.InsertNewAd("test", "2021-01-01", "2021-01-02")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if id != 1 {
		t.Errorf("expected id to be 1, but got %d", id)
	}
}

func TestInsertAgeRange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO ad_ages").
		WithArgs(1, 1, 100).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := NewAdRepository(db)
	err = r.InsertAgeRange(1, 1, 100)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestInsertCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO ad_countries").
		WithArgs(1, "US").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := NewAdRepository(db)
	err = r.InsertCountry(1, []string{"US"})
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestInsertPlatform(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO ad_platforms").
		WithArgs(1, "ios").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := NewAdRepository(db)
	err = r.InsertPlatform(1, []string{"ios"})
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestInsertGender(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO ad_genders").
		WithArgs(1, "M").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := NewAdRepository(db)
	err = r.InsertGender(1, []string{"M"})
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestGetAd(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	// rows := sqlmock.NewRows([]string{"id", "title", "start_at", "end_at"}).
	// 	AddRow(1, "test", "2021-01-01", "2021-01-02")
	// mock.ExpectQuery("SELECT id, title, start_at, end_at FROM ads WHERE id = ?").
	// 	WithArgs(1).
	// 	WillReturnRows(rows)

	// r := NewAdRepository(db)
	// ad, err := r.GetAd(1)
	// if err != nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }
	// if ad.ID != 1 {
	// 	t.Errorf("expected id to be 1, but got %d", ad.ID)
	// }
	// if ad.Title != "test" {
	// 	t.Errorf("expected title to be test, but got %s", ad.Title)
	// }
	// if ad.StartAt != "2021-01-01" {
	// 	t.Errorf("expected start_at to be 2021-01-01, but got %s", ad.StartAt)
	// }
	// if ad.EndAt != "2021-01-02" {
	// 	t.Errorf("expected end_at to be 2021-01-02, but got %s", ad.EndAt)
	// }
}
