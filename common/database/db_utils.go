package database

import (
	"log/slog"

	"github.com/lib/pq"
	"github.com/n1kos03/User_Recommendations/common/models"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (db *Database) GetAllUsers() ([]models.User, error) {
	var users []models.User

	rows, err := db.Conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			slog.Error("Error closing rows", "Error: ", err)
		}
	}()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Password, pq.Array(&user.FavoriteProduct), &user.CreatedAt, &user.UpdatedAt, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (db *Database) InsertUser(user *models.User) error {
	_, err := db.Conn.Exec("INSERT INTO users (name, email, password, favorite_product) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password, pq.Array(user.FavoriteProduct))
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetUserByID(id string) models.User {
	var user models.User
	err := db.Conn.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Password, pq.Array(&user.FavoriteProduct), &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}
	}

	return user
}

func (db *Database) UpdateUser(id string, favoriteProduct string) (models.User, error) {
	var user models.User
	err := db.Conn.QueryRow("UPDATE users SET favorite_product = $1 WHERE id = $2 RETURNING *", favoriteProduct, id).Scan(&user.ID, &user.Name, &user.Password, &user.FavoriteProduct, &user.CreatedAt, &user.UpdatedAt, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
