package domain

import (
	"time"
)

// ShortLink ...
type ShortLink struct {
	Id             string    `json:"id"`
	Code           string    `json:"code"`
	Url            string    `json:"url"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ExpiredAt      time.Time `json:"expiredAt"`
	VisitorCounter uint64    `json:"visitorCounter"`
}

func (shortlink *ShortLink) AddVisitorCount() {
	shortlink.VisitorCounter += 1
}
