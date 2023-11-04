package http

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var FiberModule = fx.Module(
	"fiberModule",
	fx.Provide(NewClient),
	fx.Invoke(HookFiber),
)

func NewClient(c Config) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: c.DisableStartupMessage,
	})

	app.Use(cors.New())

	return app
}

func HookFiber(lc fx.Lifecycle, app *fiber.App, l *zap.SugaredLogger, c Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := app.Listen(fmt.Sprintf(":%v", c.Port)); err != nil {
					l.Error(err)
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			return app.Shutdown()
		},
	})
}
