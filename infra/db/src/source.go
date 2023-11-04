package src

import (
	src "github.com/S4mkiel/sso/infra/db/src/source"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var SourceModule = fx.Module(
	"source",
	fx.Provide(NewSQLSources),
)

type Sources struct {
	UserSQL *src.UserRepository
}

func NewSQLSources(db *gorm.DB) *Sources {
	var src = Sources{
		UserSQL: src.NewUserRepository(db),
	}
	return &src
}
