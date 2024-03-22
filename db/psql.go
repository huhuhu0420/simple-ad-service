package db

import (
	"context"
	"fmt"

	"github.com/huhuhu0420/simple-ad-service/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type DB interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Ping(context.Context) error
}

func NewDB(config *utils.Config) DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if pool.Ping(context.Background()) != nil {
		logrus.Fatal(err)
		return nil
	}
	return pool
}
