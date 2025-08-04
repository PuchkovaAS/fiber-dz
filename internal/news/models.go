package news

import (
	"mime/multipart"
	"time"
)

type newsCreateForm struct {
	Title   string
	Preview *multipart.FileHeader
	Text    string
}

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
