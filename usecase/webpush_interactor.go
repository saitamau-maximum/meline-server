package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase/presenter"
	"golang.org/x/sync/errgroup"
)

type IWebPushInteractor interface {
	StoreSubscription(ctx context.Context, userID string, subscription *webpush.Subscription) error
	DeleteSubscription(ctx context.Context, key string) error
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

func (i *WebPushInteractor) StoreSubscription(ctx context.Context, userID string, subscription *webpush.Subscription) error {
	if i.webPushRepository.SetSubscription(ctx, userID, subscription) != nil {
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

func (i *WebPushInteractor) SendWebPushNotification(ctx context.Context, channelID uint64, message []byte) error {
	channel, err := i.channelRepository.FindByID(ctx, channelID)
	if err != nil {
		return fmt.Errorf("[ERROR] : failed to find channel by ID: %v", err)
	}

	users := channel.Users
	if len(users) == 0 {
		return nil
	}

	subscriptions := make([]*webpush.Subscription, 0)
	userIDs := make([]string, 0)
	for _, user := range users {
		userIDs = append(userIDs, strconv.FormatUint(user.ID, 10))
	}

	subscriptions, err = i.webPushRepository.GetSubscriptions(ctx, userIDs)
	if err != nil {
		return fmt.Errorf("[ERROR] : failed to get subscriptions: %v", err)
	}
	if len(subscriptions) == 0 {
		return nil
	}

	eg, ctx := errgroup.WithContext(ctx)

	for _, subscription := range subscriptions {
		eg.Go(func() error {
			if err := i.pushServiceRepository.SendWebPushNotification(ctx, message, subscription); err != nil {
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
