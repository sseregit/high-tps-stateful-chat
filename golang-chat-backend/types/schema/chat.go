package schema

import "time"

type Chat struct {
	ID      int64     `json:"id"`
	Room    string    `json:"room"`
	Name    string    `json:"name"`
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}

func (c *Chat) ScanRow(scanner interface{ Scan(dest ...any) error }) error {
	return scanner.Scan(
		&c.ID,
		&c.Room,
		&c.Name,
		&c.Message,
		&c.When,
	)
}
