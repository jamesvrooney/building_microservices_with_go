package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = []*Product{
	{
		ID:   1,
		Name: "car",
	},
	{
		ID:   2,
		Name: "bus",
	},
	{
		ID:   3,
		Name: "van",
	},
}

func GetProducts() Products {
	return products
}

type Products []*Product

func (p *Products) ToJson(wr io.Writer) error {
	encoder := json.NewEncoder(wr)

	return encoder.Encode(p)
}
