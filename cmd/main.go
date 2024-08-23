package main

import (
	"net/http"

	"github.com/synt4xer/go-clean-arch/cmd/app"
	"github.com/synt4xer/go-clean-arch/cmd/fx"
	"github.com/synt4xer/go-clean-arch/config"
	"github.com/synt4xer/go-clean-arch/database"
	"github.com/synt4xer/go-clean-arch/internal/module"
)

func main() {
	cfg, _ := config.ProvideConfig()

	apps := app.New(cfg)
	apps.AddProviders(config.ProvideConfig, database.NewDB, fx.AddRoutes)
	apps.AddModules(module.UsersModule)
	apps.AddServers(fx.NewHttpServer)
	apps.AddInvokers(func(*http.Server) {})

	apps.Run()
}
