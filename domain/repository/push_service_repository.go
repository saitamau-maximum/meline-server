package repository

import (
	"context"

	"github.com/SherClockHolmes/webpush-go"
)

type IPushServiceRepository interface {
	SendWebPushNotification(ctx context.Context, message []byte, subscription *webpush.Subscription) error
}
