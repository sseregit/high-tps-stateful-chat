package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang-chat-backend/config"
	"golang-chat-backend/repository/kafka"
	"golang-chat-backend/types/schema"
	"log"
	"strings"
)

type Repository struct {
	cfg   *config.Config
	db    *sql.DB
	Kafka *kafka.Kafka
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
	} else if r.Kafka, err = kafka.NewKafka(cfg); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (s *Repository) ServerSet(ip string, available bool) error {
	_, err := s.db.Exec("INSERT INTO serverInfo(`ip`, `available`) VALUES(?, ?) ON DUPLICATE KEY UPDATE `available` = VALUES(`available`)", ip, available)
	return err
}

func (s *Repository) InsertChatting(user, message, roomName string) error {
	log.Println("Insert Chatting Using WSS", "from", user, "message", message, "room", roomName)
	_, err := s.db.Exec("INSERT INTO chatting.chat(room, name, message) VALUES(?, ?, ?)", roomName, user, message)
	return err
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

	if result == nil {
		result = []T{}
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

	if err = noResut(err); err != nil {
		return nil, err
	} else {
		return nil, nil
	}

}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}

func noResut(err error) error {
	if strings.Contains(err.Error(), "sql: no rows in result set") {
		return nil
	} else {
		return err
	}
}
