package pushservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type PushServiceRepository struct{}

func NewPushServiceRepository() repository.IPushServiceRepository {
	return &PushServiceRepository{}
}

func (r *PushServiceRepository) SendWebPushNotification(ctx context.Context, message []byte, subscription *webpush.Subscription) error {
	res, err := webpush.SendNotification(message, subscription, &webpush.Options{
		VAPIDPublicKey:  config.GetEnv("VAPID_PUBLIC_KEY", ""),
		VAPIDPrivateKey: config.GetEnv("VAPID_PRIVATE_KEY", ""),
		TTL:             int(24 * time.Hour / time.Second),
	})
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("[ERROR] failed to send web push notification: %v", res.Status)
	}

	return res.Body.Close()
}
