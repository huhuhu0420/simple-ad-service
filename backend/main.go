package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/huhuhu0420/simple-ad-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_adHandler "github.com/huhuhu0420/simple-ad-service/ads/handler"
	_adRepository "github.com/huhuhu0420/simple-ad-service/ads/repository"
	_adService "github.com/huhuhu0420/simple-ad-service/ads/service"
	"github.com/huhuhu0420/simple-ad-service/utils"

	_ "github.com/lib/pq"
)

func HandlHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}
}

//	@title			Advertising ServiceAPI
//	@version		1.0
//	@description	This is a simple advertising service API
//	@host			localhost:5000
//	@BasePath		/api/v1

func main() {
	logrus.SetReportCaller(true)

	config, err := utils.LoadConfig("../")
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default(), HandlHeaders())

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adRepo := _adRepository.NewAdRepository(db)
	adService := _adService.NewAdService(adRepo)
	_adHandler.NewAdHandler(r, adService)

	logrus.Info("HTTP server started!!!")
	logrus.Fatal(r.Run(config.RestfulHost + ":" + config.RestfulPort))
}
