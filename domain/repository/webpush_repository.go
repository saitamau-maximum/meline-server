package repository

import (
	"context"

	model "github.com/saitamau-maximum/meline/models"
)

type IWebPushRepository interface {
	Create(ctx context.Context, subscription *model.Subscription) error
	FindByUserIds(ctx context.Context, subscriptions []*model.Subscription) ([]*model.Subscription, error)
	Delete(ctx context.Context, id string) error
}
