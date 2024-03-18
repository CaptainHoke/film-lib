package repo

import (
	"context"
	"errors"
	actorservice "film-lib/gen/actor_service"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	PgInstance *Postgres
	pgOnce     sync.Once
)

func NewPostgresDB(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			_ = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}

		PgInstance = &Postgres{db}
	})

	return PgInstance, nil
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

func (pg *Postgres) GetAllActors(ctx context.Context) []*actorservice.ActorWithFilms {
	query := `
		SELECT actor_id, name, sex, birthdate, array_agg(f.film_id) AS film_ids
		FROM
			actors a INNER JOIN actor_film af ON a.actor_id = af.actor_id
		GROUP BY a.name;`
	rows, _ := pg.Db.Query(ctx, query)
	defer rows.Close()

	var actors []*actorservice.ActorWithFilms
	for rows.Next() {
		a := actorservice.ActorWithFilms{}

		err := rows.Scan(&a.ActorResult.ActorID, &a.ActorResult.ActorName, &a.ActorResult.ActorSex, &a.ActorResult.ActorBirthdate, &a.FilmIds)
		if err != nil {
			fmt.Errorf("failed to scan row: %w", err)
			return nil
		}
		actors = append(actors, &a)
	}

	return actors
}

func (pg *Postgres) AddActor(ctx context.Context, info actorservice.ActorInfo) *actorservice.ActorResult {
	query := `
		INSERT INTO actors (name, sex, birthdate)
		VALUES (@actorName, @actorSex, @actorDate)
		RETURNING actor_id;`
	args := pgx.NamedArgs{
		"actorName": info.ActorName,
		"actorSex":  info.ActorSex,
		"actorDate": info.ActorBirthdate,
	}
	res := pg.Db.QueryRow(ctx, query, args)
	var actorId uint64
	_ = res.Scan(&actorId)

	return &actorservice.ActorResult{actorId, &info.ActorName, &info.ActorSex, &info.ActorBirthdate}
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
