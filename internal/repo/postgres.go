package repo

import (
	"context"
	"errors"
	actorservice "film-lib/gen/actor_service"
	searchservice "film-lib/gen/search_service"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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
	if err != nil {
		return false, errors.New("table existence query failed")
	}
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

func (pg *Postgres) GetAllActors(ctx context.Context, logger *log.Logger) []searchservice.ActorWithFilmsResult {
	//query := `
	//	SELECT actor_id, name, sex, birthdate, array_agg(f.film_id) AS film_ids
	//	FROM
	//		actors a LEFT JOIN actor_film af ON a.actor_id = af.actor_id
	//	GROUP BY actor_id;`
	query := `SELECT * FROM actors`
	rows, err := pg.Db.Query(ctx, query)

	if err != nil {
		logger.Println("failed to get all actors")
		return nil
	}

	defer rows.Close()

	var actors []searchservice.ActorWithFilmsResult
	for rows.Next() {
		a := searchservice.ActorWithFilmsResult{}

		var name string
		var sex string
		var date pgtype.Date
		err := rows.Scan(&a.ActorID, &name, &sex, &date)

		a.ActorName = &name
		a.ActorSex = &sex
		formattedDate := date.Time.Format("2006-01-02")
		a.ActorBirthdate = &formattedDate

		if err != nil {
			logger.Printf("failed to scan row: %v", err)
			return nil
		}
		actors = append(actors, a)
	}

	return actors
}

func (pg *Postgres) AddActor(ctx context.Context, info actorservice.ActorInfo) uint64 {
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
	actorId := ^uint64(0)
	_ = res.Scan(&actorId)

	return actorId
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
