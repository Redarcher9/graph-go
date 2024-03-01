package services

import (
	"Gographql/graph/model"
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

func (bi *MobileInteractor) AddMobile(ctx context.Context, input model.NewMobile) (model.Mobile, error) {
	return bi.Repo.AddMobile(ctx, input)
}

func (bi *MobileInteractor) GetMobileByName(ctx context.Context, name string) (model.Mobile, error) {
	return bi.Repo.GetMobileByName(ctx, name)
}

func (bi *MobileInteractor) UpdateMobile(ctx context.Context, input model.NewMobile) (model.Mobile, error) {
	return bi.Repo.UpdateMobile(ctx, input)
}
