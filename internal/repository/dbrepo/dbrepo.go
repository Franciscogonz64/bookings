package dbrepo

import (
	"database/sql"

	"github.com/franciscogonz64/bookings/internal/config"
	"github.com/franciscogonz64/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
