package db

import (
	"context"

	"github.com/S4mkiel/sso/domain/entity"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresModule = fx.Module(
	"postgres",
	fx.Provide(NewClient),
	fx.Invoke(HookDatabase),
	fx.Invoke(migrate),
)

func NewClient(cfg Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func HookDatabase(lc fx.Lifecycle, db *gorm.DB, logger *zap.SugaredLogger) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				dbDriver, err := db.DB()
				if err != nil {
					logger.Fatalf("failed to get db driver: %v", err)
					return nil
				}

				err = dbDriver.Ping()
				if err != nil {
					logger.Fatalf("failed to ping db: %v", err)
					return nil
				}

				err = enableUUIDExtension(db)
				if err != nil {
					logger.Fatalf("failed to enable uuid extension: %v", err)
					return nil
				}

				logger.Info("database connected")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				dbDriver, err := db.DB()
				if err != nil {
					logger.Fatalf("failed to get db driver: %v", err)
					return nil
				}

				err = dbDriver.Close()
				if err != nil {
					logger.Fatalf("failed to close db: %v", err)
					return nil
				}

				logger.Info("database disconnected")
				return nil
			},
		},
	)
}

func enableUUIDExtension(db *gorm.DB) error {
	_, err := db.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Rows()
	return err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
	)
}
