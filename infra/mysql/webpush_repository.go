package mysql

import (
	"context"
	"sync"

	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

type WebPushRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewWebPushRepository(db *bun.DB) repository.IWebPushRepository {
	return &WebPushRepository{
		db: db,
	}
}

func (r *WebPushRepository) Create(ctx context.Context, subscription *model.Subscription) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := r.db.NewInsert().Model(subscription).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *WebPushRepository) FindByUserIds(ctx context.Context, subscriptions []*model.Subscription) ([]*model.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_subscriptions := make([]*model.Subscription, 0)
	if err := r.db.NewSelect().Model(&_subscriptions).Where("user_id").Scan(ctx); err != nil {
		return nil, err
	}

	return _subscriptions, nil
}

func (r *WebPushRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.NewDelete().Model((*model.Subscription)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
