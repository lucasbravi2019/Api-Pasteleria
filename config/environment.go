package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

func LoadEnv() {
	err := godotenv.Load()

	if pkg.HasError(err) {
		log.Fatal(err)
	}
}
