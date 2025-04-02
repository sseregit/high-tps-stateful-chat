package schema

import "time"

type Room struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (r *Room) ScanRow(scanner interface{ Scan(dest ...any) error }) error {
	return scanner.Scan(
		&r.ID,
		&r.Name,
		&r.CreateAt,
		&r.UpdateAt,
	)
}
