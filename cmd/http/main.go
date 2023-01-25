package main

import (
	"log"

	"github.com/Shteyd/ddos-guard-test/config"
	"github.com/Shteyd/ddos-guard-test/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.RunHTTP(cfg)
}
