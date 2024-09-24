package blacklist

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
	tableName = "blacklists"

	columnID     = "id"
	columnSubnet = "subnet"
)

type repo struct {
	db db.Client
}

// NewRepository создание репозитория.
func NewRepository(db db.Client) repository.BlackListRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, dto model.BlackList) (int64, error) {
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
		Name:     "blacklist.Create",
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
		Name:     "blacklist.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error(err.Error())
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
		Name:     "blacklist.IsInList",
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
