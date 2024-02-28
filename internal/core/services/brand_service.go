package services

import (
	"Gographql/internal/core/domain"
	"context"
)

type BrandInteractor struct {
	Repo BrandRepo
}

func NewBrandInteractor(repo BrandRepo) *BrandInteractor {
	if repo == nil {
		return nil
	}
	return &BrandInteractor{
		Repo: repo,
	}
}

func (bi *BrandInteractor) GetBrands(ctx context.Context) ([]domain.Brand, error) {
	return bi.Repo.GetBrands(ctx)
}
