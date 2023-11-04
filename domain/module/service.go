package module

import (
	"github.com/S4mkiel/sso/domain/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Provide(service.NewUserService),
)
