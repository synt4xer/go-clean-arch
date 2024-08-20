package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/synt4xer/go-clean-arch/config"
	"github.com/synt4xer/go-clean-arch/pkg/logger"
	"go.uber.org/fx"
)

type App struct {
	Config    *config.Config
	Modules   []fx.Option
	Servers   []any
	Providers []any
	Invokers  []any
}

func New(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (a *App) AddModules(modules ...fx.Option) {
	a.Modules = append(a.Modules, modules...)
}

func (a *App) AddServers(servers ...any) {
	a.Servers = append(a.Servers, servers...)
}

func (a *App) AddProviders(providers ...any) {
	a.Providers = append(a.Providers, providers...)
}

func (a *App) AddInvokers(invokers ...any) {
	a.Invokers = append(a.Invokers, invokers...)
}

func (a *App) Run() {
	var opts = []fx.Option{
		fx.WithLogger(logger.Default()),
	}

	opts = append(opts, a.Modules...)
	opts = append(opts, fx.Provide(a.Providers...))
	opts = append(opts, fx.Provide(a.Servers...))
	opts = append(opts, fx.Invoke(a.Invokers...))

	fx.New(opts...).Run()
}
