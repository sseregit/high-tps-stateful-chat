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

func (s *Repository) GetChatList(roomName string) ([]*schema.Chat, error) {
	qs := query([]string{"SELECT * FROM", chat, "WHERE room = ? ORDER BY `when` DESC LIMIT 10"})

	rows, err := s.db.Query(qs, roomName)
	if err != nil {
		return nil, err
	}

	return getList(rows, func() *schema.Chat { return new(schema.Chat) })
}

func (s *Repository) RoomList() ([]*schema.Room, error) {
	qs := query([]string{"SELECT * FROM", room})

	rows, err := s.db.Query(qs)
	if err != nil {
		return nil, err
	}

	return getList(rows, func() *schema.Room { return new(schema.Room) })
}

func getList[T schema.Scannable](rows *sql.Rows, constructor func() T) ([]T, error) {

	defer rows.Close()

	var result []T

	for rows.Next() {
		d := constructor()

		if err := d.ScanRow(rows); err != nil {
			return nil, err
		}

		result = append(result, d)

	}

	return result, nil

}

func (s *Repository) MakeRoom(name string) error {
	_, err := s.db.Exec("INSERT INTO chatting.room(name) VALUES(?)", name)
	return err
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
