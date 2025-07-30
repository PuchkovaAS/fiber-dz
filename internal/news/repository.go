package news

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *slog.Logger
}

func NewNewsRepository(
	dbpool *pgxpool.Pool,
	customLogger *slog.Logger,
) *NewsRepository {
	return &NewsRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *NewsRepository) CountAll() int {
	query := "SELECT COUNT(*) FROM news"
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *NewsRepository) GetAll(limit, offset int) ([]News, error) {
	query := "SELECT * FROM news ORDER BY created_at LIMIT @limit OFFSET @offset"
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	news, err := pgx.CollectRows(rows, pgx.RowToStructByName[News])
	if err != nil {
		return nil, err
	}

	return news, nil
}
