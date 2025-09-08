package models

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price int      `json:"price"`
	Tags  []string `json:"tags"`
}

type ProductTags struct {
	ID         int    `json:"id"`
	Product_id int    `json:"product_id"`
	Tag        string `json:"name"`
}
