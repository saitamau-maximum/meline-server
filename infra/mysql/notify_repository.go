package mysql

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/meline/domain/repository"
	model "github.com/saitamau-maximum/meline/models"
)

type NotifyRepository struct {
	db *bun.DB
}

func NewNotifyRepository(db *bun.DB) repository.INotifyRepository {
	return &NotifyRepository{
		db: db,
	}
}

func (r *NotifyRepository) FindByUserID(ctx context.Context, userID uint64) ([]*model.Notify, error) {
	var notify []*model.Notify

	if err := r.db.NewSelect().Model(&notify).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}

	return notify, nil
}

func (r *NotifyRepository) BulkCreate(ctx context.Context, notifies []model.Notify) error {
	if _, err := r.db.NewInsert().Model(&notifies).Exec(ctx); err != nil {
		return err
	}

	return nil
}
