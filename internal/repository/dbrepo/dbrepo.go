package dbrepo

import (
	"database/sql"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// bu
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
