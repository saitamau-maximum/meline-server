package redis

import (
	"context"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/redis/go-redis/v9"
	"github.com/saitamau-maximum/meline/config"
	"github.com/saitamau-maximum/meline/domain/repository"
)

type WebPushRepository struct {
	redisClient *redis.Client
}

func NewWebPushRepository(redisClient *redis.Client) repository.IWebPushRepository {
	return &WebPushRepository{
		redisClient: redisClient,
	}
}

func (r *WebPushRepository) SetSubscription(ctx context.Context, key string, subscription map[string]*webpush.Subscription) error {
	_, err := r.redisClient.HSet(ctx, key, subscription).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *WebPushRepository) GetSubscription(ctx context.Context, key string) (*webpush.Subscription, error) {
	subscription, err := r.redisClient.HGet(ctx, key, config.SUBSCRIPTION_FIELD).Result()
	if err != nil {
		return nil, err
	}
	return &webpush.Subscription{Endpoint: subscription}, nil
}

func (r *WebPushRepository) DeleteSubscription(ctx context.Context, key string) error {
	_, err := r.redisClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

