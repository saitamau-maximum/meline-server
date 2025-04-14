package migration

import (
	"context"
	"fmt"

	model "github.com/saitamau-maximum/meline/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		// Create the subscriptions table
		if _, err := db.NewCreateTable().Model((*model.Subscription)(nil)).Exec(ctx); err != nil {
			return fmt.Errorf("failed to create subscriptions table: %w", err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		// Drop the subscriptions table
		if _, err := db.NewDropTable().Model((*model.Subscription)(nil)).IfExists().Exec(ctx); err != nil {
			return fmt.Errorf("failed to drop subscriptions table: %w", err)
		}
		return nil
	})
}
