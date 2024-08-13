package drivers

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/synt4xer/go-clean-arch/config"
	"time"
)

func NewDB(ctx context.Context, cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(95)

	return db, nil
}
