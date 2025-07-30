package news

import "time"

type News struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Preview    string    `db:"preview"`
	Text       string    `db:"text"`
	UserID     int       `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	Categories []string  `db:"categories"`
	Alias      string    `db:"alias"`
}
