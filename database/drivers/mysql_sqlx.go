package drivers

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/synt4xer/go-clean-arch/config"
	"go.uber.org/fx"
	"os"
)

func NewDB(lc fx.Lifecycle, cfg *config.Config) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sqlx.ConnectContext(context.Background(), "mysql", dsn)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to database")
		os.Exit(1)
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			_ = db.Close()
			return nil
		},
	})

	return db
}
