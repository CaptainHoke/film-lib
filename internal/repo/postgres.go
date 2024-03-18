package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPostgresDB(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			_ = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}

		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Db.Ping(ctx)
}

func (pg *Postgres) TableExists(ctx context.Context, tableName string) (bool, error) {
	exists := false

	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public'
			AND table_name = @table
		);`
	args := pgx.NamedArgs{"table": tableName}
	rows, err := pg.Db.Query(ctx, query, args)
	defer rows.Close()

	if err != nil {
		return false, errors.New("failed to check for table existence")
	}

	err = rows.Scan(&exists)

	if err != nil {
		return false, errors.New("failed to scan row while checking for table existence")
	}

	return exists, nil
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
