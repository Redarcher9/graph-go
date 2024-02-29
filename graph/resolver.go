package graph

import (
	"Gographql/graph/model"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type MobileProcessor interface {
	GetMobiles(ctx context.Context) ([]*model.Mobile, error)
	GetMobileByName(ctx context.Context, name string) (*model.Mobile, error)
}

type BrandProcessor interface {
	GetBrands(ctx context.Context) ([]*model.Brand, error)
}

type Resolver struct {
	MobileProcessor MobileProcessor
	BrandProcessor  BrandProcessor
}
