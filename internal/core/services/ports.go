package services

import (
	"Gographql/internal/core/domain"
	"context"
)

type (
	MobileRepo interface {
		GetMobiles(ctx context.Context) ([]domain.Mobile, error)
	}
	BrandRepo interface {
		GetBrands(ctx context.Context) ([]domain.Brand, error)
	}
)
