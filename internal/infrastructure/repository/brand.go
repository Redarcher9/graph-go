package repository

import (
	"Gographql/internal/core/domain"
	"Gographql/internal/infrastructure/repository/data"
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
	_, brands := data.Read()
	var res []domain.Brand

	for _, v := range brands {
		res = append(res, domain.Brand{
			Name:    v.Name,
			BrandID: v.BrandID,
		})
	}
	return res, nil
}
