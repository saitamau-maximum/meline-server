package usecase

import (
	"context"
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/google/uuid"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
	"golang.org/x/sync/errgroup"
)

type IWebPushInteractor interface {
	StoreSubscription(ctx context.Context, userID uint64, subscription *webpush.Subscription) error
	DeleteSubscription(ctx context.Context, id string) error
	SendWebPushNotification(ctx context.Context, channelID uint64, message []byte) error
	GetPublicKey() *response.GetPublicKeyResponse
}

type WebPushInteractor struct {
	channelRepository     repository.IChannelRepository
	webPushRepository     repository.IWebPushRepository
	pushServiceRepository repository.IPushServiceRepository
	webPushPresenter      presenter.IWebPushPresenter
}

func NewWebPushInteractor(channelRepository repository.IChannelRepository, webPushRepository repository.IWebPushRepository, pushServiceRepository repository.IPushServiceRepository, webPushPresenter presenter.IWebPushPresenter) IWebPushInteractor {
	return &WebPushInteractor{
		channelRepository:     channelRepository,
		webPushRepository:     webPushRepository,
		pushServiceRepository: pushServiceRepository,
		webPushPresenter:      webPushPresenter,
	}
}

func (i *WebPushInteractor) StoreSubscription(ctx context.Context, userID uint64, subscription *webpush.Subscription) error {
	subscriptionModel := &model.Subscription{
		ID:       i.generateSubscriptionID(),
		UserID:   userID,
		Endpoint: subscription.Endpoint,
		P256dh:   subscription.Keys.P256dh,
		Auth:     subscription.Keys.Auth,
	}

	if i.webPushRepository.Create(ctx, subscriptionModel) != nil {
		return fmt.Errorf("[ERROR] : failed to store subscription")
	}

	return nil
}

func (i *WebPushInteractor) DeleteSubscription(ctx context.Context, key string) error {
	if err := i.webPushRepository.Delete(ctx, key); err != nil {
		return fmt.Errorf("[ERROR] : failed to delete subscription")
	}
	return nil
}

func (i *WebPushInteractor) SendWebPushNotification(ctx context.Context, channelID uint64, message []byte) error {
	channel, err := i.channelRepository.FindByID(ctx, channelID)
	if err != nil {
		return fmt.Errorf("[ERROR] : failed to find channel by ID: %v", err)
	}

	users := channel.Users
	if len(users) == 0 {
		return nil
	}

	subscriptionConditions := make([]*model.Subscription, 0)
	for _, user := range users {
		subscriptionConditions = append(subscriptionConditions, &model.Subscription{
			UserID: user.ID,
		})
	}

	subscriptions, err := i.webPushRepository.FindByUserIds(ctx, subscriptionConditions)
	if err != nil {
		return fmt.Errorf("[ERROR] : failed to get subscriptions: %v", err)
	}
	if len(subscriptions) == 0 {
		return nil
	}

	eg, ctx := errgroup.WithContext(ctx)

	for _, subscription := range subscriptions {
		eg.Go(func() error {
			if err := i.pushServiceRepository.SendWebPushNotification(ctx, message, subscription.ToWebPushSubscription()); err != nil {
				return err
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("[ERROR] : failed to send web push notification: %v", err)
	}

	return nil
}

func (i *WebPushInteractor) GetPublicKey() *response.GetPublicKeyResponse {
	pubKey := config.GetEnv("VAPID_PUBLIC_KEY", "")
	return i.webPushPresenter.GetPublicKeyResponse(pubKey)
}

func (i *WebPushInteractor) generateSubscriptionID() string {
	return uuid.New().String()
}
