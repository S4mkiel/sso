package controller

import (
	"net/http"

	"github.com/S4mkiel/sso/domain/service"
	"github.com/S4mkiel/sso/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthController struct {
	userSerivce *service.UserService
	logger      *zap.SugaredLogger
}

func NewAuthController(userService *service.UserService, logger *zap.SugaredLogger) *AuthController {
	return &AuthController{userSerivce: userService, logger: logger}
}

func (c *AuthController) RegisterControllers(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("signup", c.SignUp)
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	response := new(dto.Base)

	response.Success = true
	response.Message = "Created"

	return ctx.Status(http.StatusOK).JSON(response)
}
