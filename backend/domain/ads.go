package domain

import _ "github.com/huhuhu0420/ads-service/docs"

type Ad struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Price   int    `json:"price"`
}
