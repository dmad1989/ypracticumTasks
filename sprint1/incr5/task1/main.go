package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	User        string `env:"USERNAME"` // укажите тег env
	Addr        string `env:"SERVER_ADDRESS"`
	shortAddres string `env:"BASE_URL"`
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current user is %s\n", cfg.User)
	fmt.Println(cfg.Addr)
	fmt.Printf("shortAddres is %s\n", cfg.shortAddres)
}
