package handler

import (
	"net/http"

	"github.com/huhuhu0420/simple-ad-service/domain"

	"github.com/gin-gonic/gin"
)

type AdHandler struct {
	Service domain.AdService
}

func NewAdHandler(e *gin.Engine, service domain.AdService) {
	handler := &AdHandler{
		Service: service,
	}

	v1 := e.Group("/api/v1")
	{
		v1.POST("/ad", handler.CreateAd)
		v1.GET("/ad", handler.GetAd)
	}
}

// CreateAd godoc
//
//	@Summary		Create a new ad
//	@Description	Create a new ad, the conditions are optional
//	@Tags			ads
//	@Accept			json
//	@Produce		json
//	@Param			ad	body	domain.AdInfo	true	"Ad object that needs to be added"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/ad [post]
func (ah *AdHandler) CreateAd(c *gin.Context) {
	var ad domain.AdInfo
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ah.Service.CreateAd(ad, ad.Conditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetAd godoc
//
//	@Summary		Get ad array by query
//	@Description	Get ad array by query
//	@Tags			ads
//	@Accept			json
//	@Produce		json
//	@Param			Offset		query		integer	false	"Offset"
//	@Param			Limit		query		integer false	"Limit"
//	@Param			Age			query		integer	false	"Age"		minimum(1)	maximum(100)
//	@Param			Gender		query		string	false	"Gender"	Enums(M, F)
//	@Param			Country		query		string	false	"Country"
//	@Param			Platform	query		string	false	"Platform"	Enums(Android, iOS, Web)
//	@Success		200			{object}	domain.AdsResponse
//	@Failure		400
//	@Failure		500
//	@Router			/ad [get]
func (ah *AdHandler) GetAd(c *gin.Context) {
	var searchAdRequest domain.SearchAdRequest
	if err := c.ShouldBindQuery(&searchAdRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adResponse, err := ah.Service.GetAd(searchAdRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adResponse)
}
