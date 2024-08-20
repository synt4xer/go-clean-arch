package main

import (
	"github.com/synt4xer/go-clean-arch/cmd/app"
	"github.com/synt4xer/go-clean-arch/cmd/fx"
	"github.com/synt4xer/go-clean-arch/config"
	"github.com/synt4xer/go-clean-arch/database/drivers"
	"net/http"
)

func main() {
	cfg, _ := config.ProvideConfig()

	apps := app.New(cfg)
	apps.AddProviders(config.ProvideConfig, drivers.NewDB, fx.AddRoutes)
	//apps.AddModules()
	apps.AddServers(fx.NewHttpServer)
	apps.AddInvokers(func(*http.Server) {})

	apps.Run()
}
