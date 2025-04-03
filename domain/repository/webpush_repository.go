package repository

import (
	"context"

	webpush "github.com/SherClockHolmes/webpush-go"
)

type IWebPushRepository interface {
	SetSubscription(ctx context.Context, key string, subscription *webpush.Subscription) error
	GetSubscription(ctx context.Context, key string) (*webpush.Subscription, error)
	GetSubscriptions(ctx context.Context, keys []string) ([]*webpush.Subscription, error)
	DeleteSubscription(ctx context.Context, key string) error
}
