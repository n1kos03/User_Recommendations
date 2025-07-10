package database

import (
	"database/sql"
	"net"
	"net/url"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// var DB *sql.DB

type Database struct {
	Conn *sql.DB
}

type DBConfig struct {
	User string
	Pass string
	Name string
	Host string
	Port string
	SSLMode string
}

func NewDatabase(cfg DBConfig) (*Database, error) {		
	dbURL := &url.URL{
		Scheme: "postgres",
		User: url.UserPassword(cfg.User, cfg.Pass),
		Host: net.JoinHostPort(cfg.Host, cfg.Port),
		Path: cfg.Name,
		RawQuery: "sslmode=" + url.PathEscape(cfg.SSLMode),
	}

	db, err := sql.Open("postgres", dbURL.String())
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}