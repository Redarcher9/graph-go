package kafkaHandler

import (
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

func (b *BrandProcessor) GetBrands(ctx context.Context) ([]domain.Brand, error) {
	res, err := b.brandInteractor.GetBrands(ctx)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}
