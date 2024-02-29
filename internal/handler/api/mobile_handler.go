package api

import (
	"Gographql/graph/model"
	"Gographql/internal/core/domain"
	"context"
	"fmt"
	"log"
)

type MobileProcessor struct {
	mobileInteractor MobileService
}

func NewMobileProcessor(mobileInteractor MobileService) *MobileProcessor {
	if mobileInteractor == nil {
		return nil
	}
	return &MobileProcessor{
		mobileInteractor: mobileInteractor,
	}
}

func (b *MobileProcessor) GetMobiles(ctx context.Context) ([]*model.Mobile, error) {
	res, err := b.mobileInteractor.GetMobiles(ctx)
	if err != nil {
		log.Println(err)
	}
	return mapMobileToModel(res), nil
}

func (b *MobileProcessor) GetMobileByName(ctx context.Context, name string) (*model.Mobile, error) {
	return &model.Mobile{}, nil
}

func mapMobileToModel(mobile []domain.Mobile) []*model.Mobile {
	Modelmobiles := []*model.Mobile{}

	for _, v := range Modelmobiles {
		fmt.Println(v)
		Modelmobiles = append(Modelmobiles, &model.Mobile{})
	}
	return Modelmobiles
}
