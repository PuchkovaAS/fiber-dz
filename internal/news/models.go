package news

import "time"

type News struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Preview    string    `db:"preview"`
	Text       string    `db:"text"`
	UserName   string    `db:"user_name"`
	CreatedAt  time.Time `db:"created_at"`
	Categories []string  `db:"categories"`
	Alias      string    `db:"alias"`
}
