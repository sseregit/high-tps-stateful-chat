package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"websocket-high-tps-chat/config"
)

type Repository struct {
	cfg *config.Config
	db  *sql.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error

	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
