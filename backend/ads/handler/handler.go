package handler

import (
	"net/http"

	"github.com/huhuhu0420/ads-service/domain"

	"github.com/gin-gonic/gin"
)

// CreateAd godoc
//
//	@Summary		Create a new ad
//	@Description	Create a new ad
//	@Tags			ads
//	@Accept			json
//	@Produce		json
//	@Param			ad	body		domain.Ad	true	"Ad object that needs to be added"
//	@Success		200	{object}	domain.Ad
//	@Router			/ads [post]
func CreateAd(c *gin.Context) {
	var ad domain.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ad)
}
