package redis

import (
	"context"
	"encoding/json"
	"fmt"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/redis/go-redis/v9"
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

func (r *WebPushRepository) SetSubscription(ctx context.Context, key string, subscription *webpush.Subscription) error {
	byteSubscription, err := json.Marshal(subscription)
	if err != nil {
		return err
	}

	_, _err := r.redisClient.HSet(ctx, key, byteSubscription).Result()
	if _err != nil {
		return _err
	}
	return nil
}

func (r *WebPushRepository) GetSubscription(ctx context.Context, key string) (*webpush.Subscription, error) {
	subscription, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var sub webpush.Subscription
	err = json.Unmarshal([]byte(subscription), &sub)
	if err != nil {
		return nil, err
	}
	if sub.Endpoint == "" {
		return nil, fmt.Errorf("[ERROR] subscription not found")
	}

	return &sub, nil
}

func (r *WebPushRepository) GetSubscriptions(ctx context.Context, keys []string) ([]*webpush.Subscription, error) {
	subscriptions := make([]*webpush.Subscription, 0)
	_, err := r.redisClient.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			subscription, err := pipe.Get(ctx, key).Result()
			if err != nil {
				return err
			}
			var sub webpush.Subscription
			err = json.Unmarshal([]byte(subscription), &sub)
			if err != nil {
				return err
			}
			if sub.Endpoint == "" {
				return nil
			}

			subscriptions = append(subscriptions, &sub)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(subscriptions) == 0 {
		return nil, fmt.Errorf("[ERROR] subscriptions not found")
	}

	return subscriptions, nil
}

func (r *WebPushRepository) DeleteSubscription(ctx context.Context, key string) error {
	_, err := r.redisClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
