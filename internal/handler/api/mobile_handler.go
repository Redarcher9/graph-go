package api

import (
	"Gographql/graph/model"
	"Gographql/internal/core/domain"
	"context"
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

func mapMobileToModel(mobiles []domain.Mobile) []*model.Mobile {
	Modelmobiles := []*model.Mobile{}

	for _, v := range mobiles {
		Modelmobiles = append(Modelmobiles, &model.Mobile{
			ModelID: v.ModelID,
			Name:    v.Name,
			Os:      v.OS,
			Country: v.Country,
			Brand: &model.Brand{
				Name:    v.Brand.Name,
				BrandID: v.Brand.BrandID,
			},
		})
	}
	return Modelmobiles
}
