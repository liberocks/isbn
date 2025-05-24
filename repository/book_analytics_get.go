package repository

import (
	"context"

	"isbn/model"
)

func (r *Repository) BookAnalyticsGet(ctx context.Context) (*model.BookAnalytics, error) {
	return &bookAnalytics, nil
}
