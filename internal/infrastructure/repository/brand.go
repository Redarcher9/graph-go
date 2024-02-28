package repository

import (
	"Gographql/internal/core/domain"
	"context"
)

type BrandRepo struct {
	database string
}

func NewBrandRepo(gorm string) BrandRepo {
	return BrandRepo{
		database: gorm,
	}
}

func (mr *BrandRepo) GetBrands(ctx context.Context) ([]domain.Brand, error) {
	res := []domain.Brand{}
	return res, nil
}
