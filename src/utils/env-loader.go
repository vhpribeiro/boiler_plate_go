package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type IEnvLoader interface {
	GetTheEnvVariable(key string) (string, error)
}

type envLoader struct{}

func (*envLoader) GetTheEnvVariable(key string) (string, error) {
	err := godotenv.Load("./configuration/.env")
	if err != nil {
		return "Erro", err
	}

	return os.Getenv(key), nil
}

func NewEnvLoader() IEnvLoader {
	return &envLoader{}
}
