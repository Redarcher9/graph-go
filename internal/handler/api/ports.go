package api

import (
	"Gographql/graph/model"
	"Gographql/internal/core/domain"
	"context"
)

type (
	MobileService interface {
		GetMobiles(ctx context.Context) ([]domain.Mobile, error)
		AddMobile(ctx context.Context, input model.NewMobile) (model.Mobile, error)
		GetMobileByName(ctx context.Context, name string) (model.Mobile, error)
		UpdateMobile(ctx context.Context, input model.NewMobile) (model.Mobile, error)
		GetMobilesByOs(ctx context.Context, name string) ([]domain.Mobile, error)
	}

	BrandService interface {
		GetBrands(ctx context.Context) ([]domain.Brand, error)
	}
)
