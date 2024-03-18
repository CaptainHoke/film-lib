package repo

import (
	"context"
	"film-lib/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPostgresDB(ctx context.Context, dbCfg config.DbConfig) (*Postgres, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%d",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.SSLMode, dbCfg.PoolMaxConns)

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
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
