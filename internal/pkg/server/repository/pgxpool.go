package repository

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxConns          = int32(90)
	defaultMinConns          = int32(0)
	defaultMaxConnLifetime   = time.Hour
	defaultMaxConnIdleTime   = time.Minute * 30
	defaultHealthCheckPeriod = time.Minute
	defaultConnectTimeout    = time.Second * 5
)

func PGXPoolConfig() *pgxpool.Config {
	//const databaseURL = "postgres://postgres:postgres@localhost:5432/o3_test_task"
	const databaseURL = "postgres://postgres:postgres@postgres:5432/o3_test_task"

	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(_ context.Context, _ *pgx.Conn) bool {
		// log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(_ *pgx.Conn) bool {
		// log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(_ *pgx.Conn) {
		// log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
