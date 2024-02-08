package handler

import (
	"net/http"

	"github.com/huhuhu0420/ads-service/domain"

	"github.com/gin-gonic/gin"
)

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
func CreateAd(c *gin.Context) {
	var ad domain.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ad)
}

// GetAd godoc
//
//	@Summary		Get ad array by query
//	@Description	Get ad array by query
//	@Tags			ads
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int		false	"Offset"
//	@Param			limit		query		int		false	"Limit"
//	@Param			age			query		int		false	"Age"		minimum(1)	maximum(100)
//	@Param			Gender		query		string	false	"Gender"	Enums(M, F)
//	@Param			Country		query		string	false	"Country"
//	@Param			Platform	query		string	false	"Platform"	Enums(Android, iOS, Web)
//	@Success		200			{object}	domain.AdsResponse
//	@Failure		400
//	@Failure		500
//	@Router			/ad [get]
func GetAd(c *gin.Context) {
	var ad domain.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ad)
}
