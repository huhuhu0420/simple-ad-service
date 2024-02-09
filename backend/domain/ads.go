package domain

import (
	_ "github.com/huhuhu0420/ads-service/docs"
)

// Ad represents an advertisement information
// @Description A brief advertisement information
type Ad struct {
	Title string `json:"title"`
	EndAt string `json:"endAt"`
}

// AdsResponse represents the response structure containing multiple ads
// @Description The response structure containing multiple ads
type AdsResponse struct {
	Items []Ad `json:"items"` // An array of ads
}

// Conditions represents the conditions for getting ads
// @Description The conditions for getting ads
type Conditions struct {
	AgeStart int    `json:"ageStart"`
	AgeEnd   int    `json:"ageEnd"`
	Country  string `json:"country"`
	Platform string `json:"platform"`
}

// AdInfo represents an advertisement information
// @Description An advertisement information
type AdInfo struct {
	Title      string     `json:"title"`
	StartAt    string     `json:"startAt"`
	EndAt      string     `json:"endAt"`
	Conditions Conditions `json:"conditions,omitempty"`
}

// SearchAdRequest represents the request structure for searching ads
// @Description The request structure for searching ads
type SearchAdRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
	Platform string `json:"platform"`
}

type AdService interface {
	CreateAd(ad AdInfo, conditions Conditions) error
	GetAd(searchAdRequest SearchAdRequest) (AdsResponse, error)
}

type AdRepository interface {
	CreateAd(ad AdInfo, conditions Conditions) error
	GetAd(searchAdRequest SearchAdRequest) (AdsResponse, error)
}
