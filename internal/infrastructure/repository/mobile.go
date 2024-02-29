package repository

import (
	"Gographql/internal/core/domain"
	"Gographql/internal/infrastructure/repository/data"
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
	mobiles, brands := data.Read()
	var res []domain.Mobile

	for _, v := range mobiles {
		var Brand domain.Brand
		for _, b := range brands {
			if b.BrandID == v.BrandID {
				Brand = domain.Brand{
					Name:    b.Name,
					BrandID: b.BrandID,
				}
			}
		}

		res = append(res, domain.Mobile{
			ModelID: v.ModelID,
			Name:    v.Name,
			OS:      v.OS,
			Country: v.Country,
			Brand:   Brand,
		})
	}
	return res, nil
}
