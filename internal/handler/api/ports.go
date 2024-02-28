package api

import (
	"Gographql/internal/core/domain"
	"context"
)

type (
	MobileService interface {
		GetMobiles(ctx context.Context) ([]domain.Mobile, error)
	}

	BrandService interface {
		GetBrands(ctx context.Context) ([]domain.Brand, error)
	}
)
