package database

import (
	"database/sql"
	"log/slog"
	"net"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

type DBConfig struct {
	User    string
	Pass    string
	Name    string
	Host    string
	Port    string
	SSLMode string
}

func NewDatabase(cfg DBConfig) (*Database, error) {
	dbURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Pass),
		Host:     net.JoinHostPort(cfg.Host, cfg.Port),
		Path:     cfg.Name,
		RawQuery: "sslmode=" + url.PathEscape(cfg.SSLMode),
	}

	slog.Info("Connecting to database", "URL: ", dbURL.String())

	db, err := sql.Open("postgres", dbURL.String())
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

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
	if err != nil && err != migrate.ErrNoChange {
		slog.Error("Error running migrations", "Error: ", err)
		return
	}

	slog.Info("Migrations successful")
}
