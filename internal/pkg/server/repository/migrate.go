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

	// migration, err := migrate.New("file://migrations", "postgres://postgres:postgres@localhost:5432/o3_test_task?sslmode=disable")
	// if err != nil {

	// 	return fmt.Errorf("something went wrong while creating new migration, err=%v", err)
	// }

	// driver, err := pgx.WithInstance(pool, &pgx.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to initialize pgx driver: %v", err)
	// }

	// m, err := migrate.NewWithDatabaseInstance("./migrations", "postgres", driver)
	// if err != nil {
	// 	log.Fatalf("Failed to create new migrate instance: %v", err)
	// }

	// if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {

	// 	return fmt.Errorf("something went wrong while up migrations, err=%v", err)
	// }

	// defer migration.Close()

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/o3_test_task?sslmode=disable")
	if err != nil {

		return fmt.Errorf("something went wrong while Open(), err=%v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {

		return fmt.Errorf("something went wrong while WithInstance(), err=%v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {

		return fmt.Errorf("something went wrong while NewWithDatabaseInstance(), err=%v", err)
	}

	_ = migration.Down()

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {

		return fmt.Errorf("something went wrong while up migrations, err=%v", err)
	}

	return nil
}
