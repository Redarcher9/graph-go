package data

import (
	"Gographql/graph/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Mobile struct {
	ModelID string `json:"modelID"`
	Name    string `json:"name"`
	OS      string `json:"os"`
	Country string `json:"country"`
	BrandID string `json:"brandID"`
}

type Brand struct {
	Name    string `json:"name"`
	BrandID string `json:"brandID"`
}

func Read() ([]Mobile, []Brand) {
	// Read the JSON file
	file, err := ioutil.ReadFile("internal/infrastructure/repository/data/data.json")
	if err != nil {
		log.Panic("Error reading JSON file:", err)
	}

	// Define struct variables to unmarshal JSON data
	var data map[string][]map[string]string
	var mobiles []Mobile
	var brands []Brand

	// Unmarshal JSON data into structs
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	// Extract mobiles and brands data from JSON
	for _, mobile := range data["mobiles"] {
		mobiles = append(mobiles, Mobile{
			ModelID: mobile["modelID"],
			Name:    mobile["name"],
			OS:      mobile["os"],
			Country: mobile["country"],
			BrandID: mobile["brandID"],
		})
	}

	for _, brand := range data["brands"] {
		brands = append(brands, Brand{
			Name:    brand["name"],
			BrandID: brand["brandID"],
		})
	}
	return mobiles, brands
}

func AddMobile(input model.NewMobile) (model.Mobile, error) {
	mobiles, brands := Read()
	_ = append(mobiles, Mobile{
		ModelID: input.ModelID,
		Name:    input.Name,
		OS:      input.Os,
		Country: input.Country,
		BrandID: input.BrandID,
	})

	var brand Brand
	for _, v := range brands {
		if v.BrandID == input.BrandID {
			brand = v
		}
	}

	return model.Mobile{
		ModelID: input.ModelID,
		Name:    input.Name,
		Os:      input.Os,
		Country: input.Country,
		Brand: &model.Brand{
			Name:    brand.BrandID,
			BrandID: brand.BrandID,
		},
	}, nil
}

func UpdateMobile(input model.NewMobile) (model.Mobile, error) {
	mobiles, brands := Read()
	var mobile Mobile
	for _, m := range mobiles {
		if m.Name == input.Name {
			mobile = m
		}
	}

	mobile.ModelID = input.ModelID
	mobile.Name = input.Name
	mobile.OS = input.Os
	mobile.Country = input.Country

	var brand Brand
	for _, v := range brands {
		if v.BrandID == input.BrandID {
			brand = v
		}
	}

	return model.Mobile{
		ModelID: mobile.BrandID,
		Name:    mobile.Name,
		Os:      mobile.OS,
		Country: mobile.Country,
		Brand: &model.Brand{
			Name:    brand.BrandID,
			BrandID: brand.BrandID,
		},
	}, nil
}

func GetMobileByName(name string) (model.Mobile, error) {
	mobiles, brands := Read()

	var mobile Mobile
	for _, m := range mobiles {
		if m.Name == name {
			mobile = m
		}
	}

	var brand Brand
	for _, v := range brands {
		if v.BrandID == mobile.BrandID {
			brand = v
		}
	}

	return model.Mobile{
		ModelID: mobile.BrandID,
		Name:    mobile.Name,
		Os:      mobile.OS,
		Country: mobile.Country,
		Brand: &model.Brand{
			Name:    brand.BrandID,
			BrandID: brand.BrandID,
		},
	}, nil
}
