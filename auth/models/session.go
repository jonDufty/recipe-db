package auth

import (
	"net"
	"time"

	"github.com/oklog/ulid/v2"
)

type Session struct {
	ID        ulid.ULID `json:"id" meddler:"id,ulid"`
	UserID    int       `json:"user_id" meddler:"user_id"`
	CreatedAt time.Time `json:"created_at" meddler:"created_at"`
	ExpiresAt time.Time `json:"expires_at" meddler:"expires_at"`
	IP        net.IP    `json:"ip" meddler:"ip,ip"`
}
