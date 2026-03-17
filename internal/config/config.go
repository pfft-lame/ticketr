package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	Dsn  string `validate:"required"`
	Port string
}

func LoadEnv(validate *validator.Validate) (Env, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := Env{
		Port: port,
		Dsn:  os.Getenv("GOOSE_DBSTRING"),
	}

	err := validate.Struct(env)
	if err != nil {
		return Env{}, err
	}

	return env, nil
}
