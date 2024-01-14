package helper

import (
	"github.com/joho/godotenv"
)

func Getenv(key, fallback string) string {
	myEnv, err := godotenv.Read()
	if err != nil {
		panic(err.Error())
	}

	value := myEnv[key]
	if len(value) == 0 {
		return fallback
	}
	return value
}
