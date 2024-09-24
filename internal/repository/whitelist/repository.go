package whitelist

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/Tel3scop/brute-force-interceptor/internal/client/db"
	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	"github.com/Tel3scop/helpers/logger"
)

const (
	tableName = "whitelists"

	columnID     = "id"
	columnSubnet = "subnet"
)

type repo struct {
	db db.Client
}

// NewRepository создание репозитория.
func NewRepository(db db.Client) repository.WhiteListRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, dto model.WhiteList) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(columnSubnet).
		Values(dto.Subnet).
		Suffix("RETURNING " + columnID).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	q := db.Query{
		Name:     "whitelist.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, subnet string) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{columnSubnet: subnet}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	q := db.Query{
		Name:     "whitelist.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *repo) IsInList(ctx context.Context, ip string) (bool, error) {
	builder := sq.Select("1").
		From(tableName).
		Where(sq.Expr(columnSubnet+" >>= ?", ip)).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	q := db.Query{
		Name:     "whitelist.IsInList",
		QueryRaw: query,
	}

	var exists int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		logger.Error(err.Error())
		return false, err
	}

	return true, nil
}
