package http

type Config struct {
	Port                  string `env:"HTTP_PORT" required:"true"`
	DisableStartupMessage bool   `env:"DISABLE_STARTUP_MESSAGE"`
}
