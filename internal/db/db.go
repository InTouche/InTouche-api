package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"intouche-back-core/internal/config"
)

var (
	//go:embed migrations/*.sql
	fs embed.FS
)

func NewConnection(cfg *config.DB) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	for i := 0; i < 10; i++ {
		if err = db.Ping(); err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		break
	}
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	if err = migrateUp(cfg.URL); err != nil {
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}

	return db, nil
}

func migrateUp(url string) error {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", sourceInstance, url)
	if err != nil {
		return fmt.Errorf("failed to create new migrate instance: %w", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations up: %w", err)
	}

	return nil
}

const (
	postgresUser   = "postgres"
	schema         = "postgres"
	hostAuthMethod = "trust"
	port           = "5432/tcp"
	dbURLf         = "postgres://%s@localhost:%s/%s?sslmode=disable"
	postgresDriver = "postgres"
)

func NewTestConnection(cfg *config.DB) (*sql.DB, error) {
	ctx := context.Background()

	var env = map[string]string{
		"POSTGRES_USER":             postgresUser,
		"POSTGRES_DB":               schema,
		"POSTGRES_HOST_AUTH_METHOD": hostAuthMethod,
	}

	dbURL := func(port nat.Port) string {
		return fmt.Sprintf(dbURLf, postgresUser, port.Port(), schema)
	}

	var req = testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{port},
		Env:          env,
		WaitingFor:   wait.ForSQL(port, postgresDriver, dbURL).Timeout(time.Second * 15),
	}

	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	pgPort, err := pgContainer.MappedPort(ctx, port)
	if err != nil {
		return nil, err
	}

	cfg.URL = dbURL(pgPort)

	return NewConnection(cfg)
}
