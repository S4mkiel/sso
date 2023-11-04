package controller

import (
	"net/http"

	"github.com/S4mkiel/sso/domain/entity"
	"github.com/S4mkiel/sso/domain/service"
	"github.com/S4mkiel/sso/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

	type Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	body := new(Payload)

	if err := ctx.BodyParser(body); err != nil {
		c.logger.Error(err)
		response.Message = "Invalid payload"
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	//password for md5
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Error(err)
		response.Message = "Invalid payload"
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	//externalid for uuid

	_, uErr := c.userSerivce.Create(&entity.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Username: body.Username,
	})

	if uErr != nil {
		c.logger.Error(err)
		response.Message = "User already exists"
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Message = "Created"

	return ctx.Status(http.StatusCreated).JSON(response)
}
