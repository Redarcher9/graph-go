package kafkaHandler

import (
	"Gographql/internal/core/domain"
	"context"
)

type (
	BrandService interface {
		GetBrands(ctx context.Context) ([]domain.Brand, error)
	}
)
