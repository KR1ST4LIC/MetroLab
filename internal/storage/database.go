package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"

	"MetroLab/db"
	"MetroLab/internal/config"
)

const (
	dbDriver = "postgres"
)

func UploadDataBase(ctx context.Context, dbCfg *config.DBConnConfig) (*pgxpool.Pool, error) {
	if dbCfg.MigrationsEnable {
		err := RunMigrations(dbCfg)
		if err != nil {
			return nil, err
		}
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
	)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create conn pool")
	}

	return pool, nil
}

func RunMigrations(dbCfg *config.DBConnConfig) error {
	dataBase, err := sql.Open(dbDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.DBName,
	))
	if err != nil {
		return errors.Wrap(err, "open database")
	}

	goose.SetBaseFS(db.EmbedMigrations)

	if err = goose.SetDialect("postgres"); err != nil {
		return errors.Wrap(err, "set dialect")
	}

	if err = goose.Up(dataBase, "migrations"); err != nil {
		return errors.Wrap(err, "up migrations")
	}
	if err = dataBase.Close(); err != nil {
		return errors.Wrap(err, "close database")
	}

	return nil
}
