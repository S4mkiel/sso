package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	service "github.com/S4mkiel/sso/domain/module"
	"github.com/S4mkiel/sso/infra/config"
	"github.com/S4mkiel/sso/infra/db"
	"github.com/S4mkiel/sso/infra/http"
	"github.com/S4mkiel/sso/infra/log"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if os.Getenv("ENV") != "production" {
		loadConfig()
	}
	fx.New(
		log.Module,
		config.Module,
		service.Module,
		db.Module,
		http.Module,
	).Run()

}

func loadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
