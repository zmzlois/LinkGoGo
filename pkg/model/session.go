package model

import (
	"time"

	"github.com/zmzlois/LinkGoGo/ent"
)

type Session struct {
	Session *ent.Session
}

// Expired returns true if the session expired
func (m *Session) Expired() bool {
	return time.Now().UTC().After(m.Session.ExpiresAt.UTC())
}
