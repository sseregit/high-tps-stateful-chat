package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"websocket-high-tps-chat/config"
	"websocket-high-tps-chat/types/schema"
)

type Repository struct {
	cfg *config.Config
	db  *sql.DB
}

const (
	room       = "chatting.room"
	chat       = "chatting.chat"
	serverInfo = "chatting.serverInfo"
)

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error

	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (s *Repository) Room(name string) (*schema.Room, error) {
	d := new(schema.Room)
	qs := query([]string{"SELECT * FROM", room, "WHERE name = ?"})

	err := s.db.QueryRow(qs, name).Scan(
		&d.ID,
		&d.Name,
		&d.CreateAt,
		&d.UpdateAt,
	)

	return d, err
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
