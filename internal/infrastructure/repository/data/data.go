package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
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
