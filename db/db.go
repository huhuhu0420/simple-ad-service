package db

import (
	"context"
	"fmt"

	"github.com/huhuhu0420/simple-ad-service/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type DB struct {
	*pgxpool.Pool
}

func NewDB(config *utils.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &DB{pool}, nil
}
