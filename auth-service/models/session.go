package models

import "time"

type Session struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	SessionTokenHash string    `json:"-"`
	ExpiresAt        time.Time `json:"expires_at"`
	Cached           bool
	RevokedAt        *time.Time `json:"revoked_at,omitempty"`
}
