package database

import (
	"User_Recommendations/internal/models"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (db *Database) CheckDBConnection() {
	err := db.Conn.Ping()
	if err != nil {
		slog.Error("Error pinging database", "Error: ", err)
	} else {
		slog.Info("Database connection established")
	}
}

func (db *Database) RunMigrations() {
	slog.Info("Running migrations")

	driver, err := postgres.WithInstance(db.Conn, &postgres.Config{})
	if err != nil {
		slog.Error("Error creating postgres driver", "Error: ", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"user_db", driver)
	if err != nil {
		slog.Error("Error creating migrate instance", "Error: ", err)
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange{
		slog.Error("Error running migrations", "Error: ", err)
		return
	}

	slog.Info("Migrations successful")
}

func (db *Database) GetAllUsers() ([]models.User, error) {
	var users []models.User

	rows, err := db.Conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.FavoriteProduct, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (db *Database) InsertUser(user *models.User) error {
	_, err := db.Conn.Exec("INSERT INTO users (name, password, favorite_product) VALUES ($1, $2, $3)", user.Name, user.Password, user.FavoriteProduct)
	if err != nil {
		return err
	}

	return nil
}