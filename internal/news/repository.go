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

type SearchParam struct {
	Limit    int
	Offset   int
	Category string
	Keyword  string
}

func (r *NewsRepository) GetByParam(param SearchParam) ([]News, error) {
	query := `
        SELECT * FROM news 
        WHERE $1 = ANY(categories)
        ORDER BY created_at DESC 
        LIMIT $2 OFFSET $3
    `

	rows, err := r.Dbpool.Query(
		context.Background(),
		query,
		param.Category, // $1
		param.Limit,    // $2
		param.Offset,   // $3
	)
	if err != nil {
		return nil, err
	}

	news, err := pgx.CollectRows(rows, pgx.RowToStructByName[News])
	if err != nil {
		return nil, err
	}

	return news, nil
}
