package http

import (
	"github.com/S4mkiel/sso/infra/http/controller"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	FiberModule,
	fx.Provide(controller.NewAuthController),
	fx.Invoke(RegisterControllers),
)

func RegisterControllers(app *fiber.App, authController *controller.AuthController) {
	v1 := app.Group("v1")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 👋")
	})

	authController.RegisterControllers(v1)
}
