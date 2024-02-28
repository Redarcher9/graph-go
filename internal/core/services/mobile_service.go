package services

import (
	"Gographql/internal/core/domain"
	"context"
)

type MobileInteractor struct {
	Repo MobileRepo
}

func NewMobileInteractor(repo MobileRepo) *MobileInteractor {
	if repo == nil {
		return nil
	}
	return &MobileInteractor{
		Repo: repo,
	}
}

func (bi *MobileInteractor) GetMobiles(ctx context.Context) ([]domain.Mobile, error) {
	return bi.Repo.GetMobiles(ctx)
}
