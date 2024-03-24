package config

import "github.com/joho/godotenv"

func GetEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Fatal err, ENV file")
	}
}
