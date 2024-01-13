package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	User string `env:"USERNAME"` // укажите тег env
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current user is %s\n", cfg.User)
}
