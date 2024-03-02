package domain

import (
	_ "github.com/huhuhu0420/simple-ad-service/docs"
)

// Ad represents an advertisement information
// @Description A brief advertisement information
type Ad struct {
	Title string `json:"title" example:"Ad title"`
	EndAt string `json:"endAt" example:"2023-12-01T00:00:00Z" format:"date-time"`
}

// AdsResponse represents the response structure containing multiple ads
// @Description The response structure containing multiple ads
type AdsResponse struct {
	Items []Ad `json:"items"` // An array of ads
}

// Conditions represents the conditions for getting ads
// @Description The conditions for getting ads
type Conditions struct {
	AgeStart int      `json:"ageStart" example:"18"`
	AgeEnd   int      `json:"ageEnd" example:"65"`
	Country  []string `json:"country" example:"US"`
	Platform []string `json:"platform" example:"Android"`
	Gender   []string `json:"gender" example:"M"`
}

// AdInfo represents an advertisement information
// @Description An advertisement information
type AdInfo struct {
	Title      string     `json:"title" binding:"required" example:"Ad title"`
	StartAt    string     `json:"startAt" binding:"required" example:"2023-01-01T00:00:00Z" format:"date-time"`
	EndAt      string     `json:"endAt" binding:"required" example:"2023-12-01T00:00:00Z" format:"date-time"`
	Conditions Conditions `json:"conditions,omitempty"`
}

// SearchAdRequest represents the request structure for searching ads
// @Description The request structure for searching ads
type SearchAdRequest struct {
	Offset   int    `json:"offset" example:"0"`
	Limit    int    `json:"limit" example:"10"`
	Age      int    `json:"age" example:"18"`
	Gender   string `json:"gender" example:"M"`
	Country  string `json:"country" example:"US"`
	Platform string `json:"platform" example:"Android"`
}

type AdService interface {
	CreateAd(ad AdInfo, conditions Conditions) error
	GetAd(searchAdRequest SearchAdRequest) (*AdsResponse, error)
}

type AdRepository interface {
	InsertNewAd(title string, startAt string, endAt string) (int, error)
	InsertAgeRange(id int, ageStart int, ageEnd int) error
	InsertCountry(id int, country []string) error
	InsertPlatform(id int, platform []string) error
	InsertGender(id int, gender []string) error
	GetAd(searchAdRequest SearchAdRequest) (*AdsResponse, error)
}
