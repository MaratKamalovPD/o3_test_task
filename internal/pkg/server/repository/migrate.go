package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MakeMigrations(migrationURL string, dbSource string) error {

	db, err := sql.Open("postgres", dbSource)
	if err != nil {

		return fmt.Errorf("something went wrong while Open(), err=%v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {

		return fmt.Errorf("something went wrong while WithInstance(), err=%v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {

		return fmt.Errorf("something went wrong while NewWithDatabaseInstance(), err=%v", err)
	}

	_ = migration.Down()

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {

		return fmt.Errorf("something went wrong while up migrations, err=%v", err)
	}

	defer migration.Close()

	return nil
}
