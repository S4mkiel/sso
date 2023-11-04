package db

import "fmt"

type Config struct {
	Host     string `env:"POSTGRES_HOST,required=true"`
	User     string `env:"POSTGRES_USER,required=true"`
	Password string `env:"POSTGRES_PASSWORD,required=true"`
	DB       string `env:"POSTGRES_DB,required=true"`
	Port     string `env:"POSTGRES_PORT,required=true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		c.Host, c.User, c.Password, c.DB, c.Port,
	)
}
