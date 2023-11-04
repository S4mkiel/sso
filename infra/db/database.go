package db

import (
	"github.com/S4mkiel/sso/domain/repository"
	"github.com/S4mkiel/sso/infra/db/src"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	PostgresModule,
	src.SourceModule,
	fx.Provide(func(src *src.Sources) repository.UserRepository { return src.UserSQL }),
)
