package usecase_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/saitamau-maximum/meline/generated/proto/go/schema/response"
	"github.com/saitamau-maximum/meline/usecase"
	"github.com/stretchr/testify/assert"

	"github.com/SherClockHolmes/webpush-go"
)

func TestWebPushInteractor_Success_StoreSubscription(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)

	sub := &webpush.Subscription{
		Endpoint: "https://example.com/endpoint",
		Keys: webpush.Keys{
			Auth:   "auth-key",
			P256dh: "p256dh-key",
		},
	}

	err := webPushInteractor.StoreSubscription(ctx, "test-key", sub)
	assert.NoError(t, err)
}

func TestWebPushInteractor_Failed_StoreSubscription(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)

	sub := &webpush.Subscription{
		Endpoint: "https://example.com/endpoint",
		Keys: webpush.Keys{
			Auth:   "auth-key",
			P256dh: "p256dh-key",
		},
	}

	ctx = context.WithValue(ctx, "is_set_test_fail", true)
	err := webPushInteractor.StoreSubscription(ctx, "test-key", sub)
	assert.Error(t, err)
}

func TestWebPushInteractor_Success_DeleteSubscription(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)

	err := webPushInteractor.DeleteSubscription(ctx, "test-key")
	assert.NoError(t, err)
}

func TestWebPushInteractor_Failed_DeleteSubscription(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)
	ctx = context.WithValue(ctx, "is_delete_test_fail", true)
	err := webPushInteractor.DeleteSubscription(ctx, "test-key")
	assert.Error(t, err)
}

func TestWebPushInteractor_Success_SendWebPushNotification(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)

	err := webPushInteractor.SendWebPushNotification(ctx, 1, []byte("test message"))
	assert.NoError(t, err)
}

func TestWebPushInteractor_Failed_SendWebPushNotification(t *testing.T) {
	ctx := context.Background()
	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)
	ctx = context.WithValue(ctx, "is_send_test_fail", true)
	err := webPushInteractor.SendWebPushNotification(ctx, 1, []byte("test message"))
	assert.Error(t, err)
}

func TestWebPushInteractor_Success_GetPublicKey(t *testing.T) {
	// envのVAPID_PUBLIC_KEYをmockする
	os.Setenv("VAPID_PUBLIC_KEY", "test-public-key")

	webPushRepo := &mockWebPushRepository{}
	channelRepo := &mockChannelRepository{}
	pushServiceRepo := &mockPushServiceRepository{}
	webPushPresenter := &mockWebPushPresenter{}
	webPushInteractor := usecase.NewWebPushInteractor(channelRepo, webPushRepo, pushServiceRepo, webPushPresenter)

	res := webPushInteractor.GetPublicKey()
	assert.Equal(t, "test-public-key", res.PublicKey)
}

type mockWebPushRepository struct{}

func (m *mockWebPushRepository) SetSubscription(ctx context.Context, key string, subscription *webpush.Subscription) error {
	if ctx.Value("is_set_test_fail") != nil {
		return fmt.Errorf("failed to set subscription")
	}

	return nil
}

func (m *mockWebPushRepository) GetSubscription(ctx context.Context, key string) (*webpush.Subscription, error) {
	if ctx.Value("is_get_test_fail") != nil {
		return nil, fmt.Errorf("failed to get subscription")
	}

	return &webpush.Subscription{
		Endpoint: "https://example.com/endpoint",
		Keys: webpush.Keys{
			Auth:   "auth-key",
			P256dh: "p256dh-key",
		},
	}, nil
}

func (m *mockWebPushRepository) GetSubscriptions(ctx context.Context, keys []string) ([]*webpush.Subscription, error) {
	if ctx.Value("is_get_test_fail") != nil {
		return nil, fmt.Errorf("failed to get subscriptions")
	}

	subscriptions := make([]*webpush.Subscription, len(keys))
	for range keys {
		subscriptions = append(subscriptions, &webpush.Subscription{
			Endpoint: "https://example.com/endpoint",
			Keys: webpush.Keys{
				Auth:   "auth-key",
				P256dh: "p256dh-key",
			},
		})
	}

	return subscriptions, nil
}

func (m *mockWebPushRepository) DeleteSubscription(ctx context.Context, key string) error {
	if ctx.Value("is_delete_test_fail") != nil {
		return fmt.Errorf("failed to delete subscription")
	}

	return nil
}

type mockWebPushPresenter struct{}

func (m *mockWebPushPresenter) GetPublicKeyResponse(pubKey string) *response.GetPublicKeyResponse {
	return &response.GetPublicKeyResponse{
		PublicKey: pubKey,
	}
}

type mockPushServiceRepository struct{}

func (m *mockPushServiceRepository) SendWebPushNotification(ctx context.Context, message []byte, subscription *webpush.Subscription) error {
	if ctx.Value("is_send_test_fail") != nil {
		return fmt.Errorf("failed to send notification")
	}

	return nil
}
