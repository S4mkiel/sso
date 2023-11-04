package config

import (
	"github.com/Netflix/go-env"
	"github.com/S4mkiel/sso/infra/db"
	"github.com/S4mkiel/sso/infra/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"config",
	fx.Provide(NewConfig),
	fx.Provide(func(cfg Config) db.Config { return cfg.Postgres }),
	fx.Provide(func(cfg Config) http.Config { return cfg.Http }),
)

type Config struct {
	Postgres db.Config
	Http     http.Config
	Extras   *env.EnvSet
}

func NewConfig(logger *zap.SugaredLogger) Config {
	var cfg Config
	err := cfg.loadConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
	return cfg

}

func (c *Config) loadConfig() error {
	es, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return err
	}
	c.Extras = &es

	return nil
}
