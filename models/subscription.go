package model

import (
	"time"

	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	ID        string    `bun:"id,pk"`
	UserID    uint64    `bun:"user_id,notnull"`
	Endpoint  string    `bun:"endpoint,notnull"`
	P256dh    string    `bun:"p256dh,notnull"`
	Auth      string    `bun:"auth,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

func (s *Subscription) ToWebPushSubscription() *webpush.Subscription {
	return &webpush.Subscription{
		Endpoint: s.Endpoint,
		Keys: webpush.Keys{
			P256dh: s.P256dh,
			Auth:   s.Auth,
		},
	}
}
