package module

import (
	"github.com/synt4xer/go-clean-arch/internal/delivery/http"
	"github.com/synt4xer/go-clean-arch/internal/repository"
	"github.com/synt4xer/go-clean-arch/internal/usecase"
	"go.uber.org/fx"
)

var UsersModule = fx.Module("users", fx.Provide(http.NewUsers), fx.Provide(usecase.New), fx.Provide(repository.New))
