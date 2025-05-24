package repository

import (
	"context"

	"isbn/model"
)

func (r *Repository) BookAnalyticsUpdate(ctx context.Context, analytics model.BookAnalytics) (*model.BookAnalytics, error) {
	bookAnalytics = analytics

	return &bookAnalytics, nil
}
