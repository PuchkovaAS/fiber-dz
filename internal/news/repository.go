package news

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

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

type SearchParam struct {
	Limit    int
	Offset   int
	Category string
	Keyword  string
}

func (r *NewsRepository) CountByParam(param SearchParam) int {
	query := "SELECT COUNT(*) FROM news"
	args := []any{}
	conditions := []string{}
	argPos := 1

	if param.Category != "" {
		conditions = append(
			conditions,
			fmt.Sprintf("$%d = ANY(categories)", argPos),
		)
		args = append(args, param.Category)
		argPos++
	}

	if param.Keyword != "" {
		conditions = append(
			conditions,
			fmt.Sprintf(
				"(title ILIKE $%d OR text ILIKE $%d)",
				argPos,
				argPos+1,
			),
		)
		args = append(args, "%"+param.Keyword+"%", "%"+param.Keyword+"%")
		argPos += 2
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var count int
	err := r.Dbpool.QueryRow(context.Background(), query, args...).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func (r *NewsRepository) CountAll() int {
	query := "SELECT COUNT(*) FROM news"
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *NewsRepository) GetAll(limit, offset int) ([]News, error) {
	query := `
        SELECT 
            n.id,
            n.title,
            n.preview,
            n.text,
            u.name AS user_name,  
            n.created_at,
            n.categories,
            n.alias
        FROM news n
        JOIN users u ON n.user_id = u.id  
        ORDER BY n.created_at DESC 
        LIMIT @limit OFFSET @offset
    `
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

func (r *NewsRepository) GetByParam(param SearchParam) ([]News, error) {
	query := `
        SELECT 
            n.id,
            n.title,
            n.preview,
            n.text,
            u.name AS user_name,
            n.created_at,
            n.categories,
            n.alias
        FROM news n
        JOIN users u ON n.user_id = u.id
    `

	args := []any{}
	conditions := []string{}
	argPos := 1

	if param.Category != "" {
		conditions = append(
			conditions,
			fmt.Sprintf("$%d = ANY(n.categories)", argPos),
		)
		args = append(args, param.Category)
		argPos++
	}

	if param.Keyword != "" {
		conditions = append(
			conditions,
			fmt.Sprintf(
				"(n.title ILIKE $%d OR n.text ILIKE $%d)",
				argPos,
				argPos+1,
			),
		)
		args = append(args, "%"+param.Keyword+"%", "%"+param.Keyword+"%")
		argPos += 2
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += `
        ORDER BY n.created_at DESC 
        LIMIT $` + strconv.Itoa(argPos) + ` OFFSET $` + strconv.Itoa(argPos+1)

	args = append(args, param.Limit, param.Offset)

	rows, err := r.Dbpool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	news, err := pgx.CollectRows(rows, pgx.RowToStructByName[News])
	if err != nil {
		return nil, fmt.Errorf("rows collection error: %w", err)
	}

	return news, nil
}
