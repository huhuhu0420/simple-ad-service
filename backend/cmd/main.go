package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func HandlHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}
}

func main() {
	logrus.SetReportCaller(true)

	logrus.Info("HTTP server started!!!")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPassword := viper.GetString("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	// if go not run in docker, host will be localhost
	if dbHost == "" {
		dbHost = "localhost"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)

	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default(), HandlHeaders())

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
