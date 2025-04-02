package usecase

import (
	"context"
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IWebPushInteractor interface {
	StoreSubscription(ctx context.Context, userID string, subscription *webpush.Subscription) error
	DeleteSubscription(ctx context.Context, key string) error
	SendWebPushNotification(ctx context.Context, message []byte, subscription *webpush.Subscription) error
	GetPublicKey() *response.GetPublicKeyResponse
}

type WebPushInteractor struct {
	webPushRepository     repository.IWebPushRepository
	pushServiceRepository repository.IPushServiceRepository
	webPushPresenter      presenter.IWebPushPresenter
}

func NewWebPushInteractor(webPushRepository repository.IWebPushRepository, pushServiceRepository repository.IPushServiceRepository, webPushPresenter presenter.IWebPushPresenter) IWebPushInteractor {
	return &WebPushInteractor{
		webPushRepository:     webPushRepository,
		pushServiceRepository: pushServiceRepository,
		webPushPresenter:      webPushPresenter,
	}
}

func (i *WebPushInteractor) StoreSubscription(ctx context.Context, userID string, subscription *webpush.Subscription) error {
	subscriptionData := map[string]*webpush.Subscription{
		config.SUBSCRIPTION_FIELD: subscription,
	}

	if i.webPushRepository.SetSubscription(ctx, userID, subscriptionData) != nil {
		return fmt.Errorf("[ERROR] : failed to store subscription")
	}

	return nil
}

func (i *WebPushInteractor) DeleteSubscription(ctx context.Context, key string) error {
	if err := i.webPushRepository.DeleteSubscription(ctx, key); err != nil {
		return fmt.Errorf("[ERROR] : failed to delete subscription")
	}
	return nil
}

func (i *WebPushInteractor) SendWebPushNotification(ctx context.Context, message []byte, subscription *webpush.Subscription) error {
	if err := i.pushServiceRepository.SendWebPushNotification(ctx, message, subscription); err != nil {
		return err
	}
	return nil
}

func (i *WebPushInteractor) GetPublicKey() *response.GetPublicKeyResponse {
	pubKey := config.GetEnv("VAPID_PUBLIC_KEY", "")
	return i.webPushPresenter.GetPublicKeyResponse(pubKey)
}
