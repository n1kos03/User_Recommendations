CREATE TABLE product_tags (
  id SERIAL PRIMARY KEY,
  product_id INT REFERENCES products(id),
  tag VARCHAR(50) NOT NULL
);