package database

import (
	"github.com/lib/pq"
	"github.com/n1kos03/User_Recommendations/common/models"
)

func (db *Database) InsertProduct(product *models.Product) error {
	err := db.Conn.QueryRow("INSERT INTO products (name, price, tags) VALUES ($1, $2, $3) RETURNING id", product.Name, product.Price, pq.Array(product.Tags)).Scan(&product.ID)
	if err != nil {
		return err
	}

	for _, tag := range product.Tags {
		_, err := db.Conn.Exec("INSERT INTO product_tags (product_id, tag) VALUES ($1, $2)", product.ID, tag)
		if err != nil {
			return err
		}
	}

	return nil
}
