package handler

import (
	"errors"
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
	err := validateCondition(ad.Conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ah.Service.CreateAd(ad, ad.Conditions)
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
	if err := c.BindQuery(&searchAdRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := validateQuery(&searchAdRequest)
	if err != nil {
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

func validateQuery(searchAdRequest *domain.SearchAdRequest) error {
	if searchAdRequest.Limit > 100 {
		return errors.New("limit should not exceed 100")
	}
	if searchAdRequest.Gender != "" && searchAdRequest.Gender != "M" && searchAdRequest.Gender != "F" {
		return errors.New("gender should be M or F")
	}
	return nil
}

func validateCondition(conditions domain.Conditions) error {
	if conditions.AgeStart > conditions.AgeEnd {
		return errors.New("ageStart should not exceed ageEnd")
	}
	if len(conditions.Country) > 0 {
		for _, c := range conditions.Country {
			if len(c) != 2 {
				return errors.New("country code should be 2 characters")
			}
		}
	}
	if len(conditions.Platform) > 0 {
		for _, p := range conditions.Platform {
			if p != "Android" && p != "iOS" && p != "Web" {
				return errors.New("platform should be Android, iOS or Web")
			}
		}
	}
	if len(conditions.Gender) > 0 {
		for _, g := range conditions.Gender {
			if g != "M" && g != "F" {
				return errors.New("gender should be M or F")
			}
		}
	}
	return nil
}
