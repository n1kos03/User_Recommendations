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

func (db *Database) GetProductByID(productID int) (models.Product, error) {
	row := db.Conn.QueryRow("SELECT * FROM products WHERE id = $1", productID)
	
	var product models.Product
	
	err := row.Scan(&product.ID, &product.Name, &product.Price, pq.Array(&product.Tags))
	if err != nil {
		return product, err
	}
	
	return product, nil
}

func (db *Database) GetProductsByTag(tags []string) ([]models.Product, error) {
	rows, err := db.Conn.Query("SELECT * FROM products WHERE $1 && tags", pq.Array(tags))
	if err != nil {
		return nil, err
	}

	var products []models.Product

	for rows.Next() {
		var product models.Product
		
		err := rows.Scan(&product.ID, &product.Name, &product.Price, pq.Array(&product.Tags))
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *Database) UpdateProductPrice(id, price int) (models.Product, error) {
	var product models.Product
	err := db.Conn.QueryRow("UPDATE products SET price = $1 WHERE id = $2 RETURNING *", price, id).Scan(&product.ID, &product.Name, &product.Price, pq.Array(&product.Tags))
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}