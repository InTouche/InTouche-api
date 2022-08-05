package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"intouche-back-core/internal/model"
)

type (
	Config struct {
		API       API       `json:"api"`
		HealthAPI HealthAPI `json:"health_api"`
		DB        DB        `json:"db"`
	}

	API struct {
		Address         string         `json:"address"`
		ReadTimeout     model.Duration `json:"read_timeout"`
		WriteTimeout    model.Duration `json:"write_timeout"`
		ShutdownTimeout model.Duration `json:"shutdown_timeout"`
	}

	HealthAPI struct {
		Port     int    `json:"port"`
		Endpoint string `json:"endpoint"`
	}

	DB struct {
		URL      string `json:"url"`
		Name     string `json:"name"`
		Host     string `json:"host"`
		Password string `json:"password"`
		Port     string `json:"port"`
		SSLMode  string `json:"ssl_mode"`
		User     string `json:"user"`

		MaxOpenConns int `json:"max_open_conns"`
		MaxIdleConns int `json:"max_idle_conns"`

		MigrateDown bool `json:"migrate_down"`
	}
)

func NewConfig(fn string) (*Config, error) {
	cfg := &Config{}

	var f io.ReadCloser
	f, err := os.Open(fn)
	if err != nil {
		return cfg, os.ErrNotExist
	}

	err = json.NewDecoder(f).Decode(cfg)
	if err != nil {
		return cfg, fmt.Errorf("unmarshalling: %s", err)
	}

	cfg.DB.URL = cfg.DB.GetURL()

	return cfg, nil
}

func (db *DB) GetURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.User, db.Password, db.Host, db.Port, db.Name, db.SSLMode,
	)
}
