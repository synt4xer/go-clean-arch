package fx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/synt4xer/go-clean-arch/config"
	"go.uber.org/fx"
	"net/http"
	"os"
	"time"
)

func NewHttpServer(lc fx.Lifecycle, router *chi.Mux, cfg *config.Config) *http.Server {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Addr:              addr,
		Handler:           router,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Err(err).Msg("error starting http server")
					os.Exit(1)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			toCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(toCtx); err != nil {
				log.Error().Err(err).Msg("can't shutdown http server")
				return err
			}
			log.Info().Msg("http server shutdown")
			return nil
		},
	})

	return srv
}
