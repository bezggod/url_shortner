package model

import "time"

type ShortURL struct { // model url
	ID          int64
	OriginalUrl string
	ShortUrl    string // max 10
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewShortURL(oURL, sURL string, now time.Time) *ShortURL {
	return &ShortURL{
		OriginalUrl: oURL,
		ShortUrl:    sURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
