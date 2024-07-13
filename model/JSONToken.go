package model

import (
	"time"
)

type JSONToken struct {
	Issuer     string    `json:"iss"`
	Subject    string    `json:"sub"`
	Audience   string    `json:"aud"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
}