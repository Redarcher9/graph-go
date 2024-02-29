package api

import (
	"Gographql/graph/model"
	"Gographql/internal/core/domain"
	"context"
	"log"
)

type BrandProcessor struct {
	brandInteractor BrandService
}

func NewBrandProcessor(brandInteractor BrandService) *BrandProcessor {
	if brandInteractor == nil {
		return nil
	}
	return &BrandProcessor{
		brandInteractor: brandInteractor,
	}
}

func (b *BrandProcessor) GetBrands(ctx context.Context) ([]*model.Brand, error) {
	res, err := b.brandInteractor.GetBrands(ctx)
	if err != nil {
		log.Println(err)
	}
	return mapBrandToModel(res), nil
}

func mapBrandToModel(brands []domain.Brand) []*model.Brand {
	Modelbrands := []*model.Brand{}

	for _, v := range brands {
		Modelbrands = append(Modelbrands, &model.Brand{
			Name:    v.Name,
			BrandID: v.BrandID,
		})
	}
	return Modelbrands
}
