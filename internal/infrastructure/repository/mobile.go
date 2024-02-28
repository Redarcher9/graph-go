package repository

import (
	"Gographql/internal/core/domain"
	"context"
)

type MobileRepo struct {
	database string
}

func NewMobileRepo(gorm string) MobileRepo {
	return MobileRepo{
		database: gorm,
	}
}

func (mr *MobileRepo) GetMobiles(ctx context.Context) ([]domain.Mobile, error) {
	res := []domain.Mobile{}
	return res, nil
}
