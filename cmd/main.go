package main

import (
	"cors/api"
	"cors/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yuklashda xato: %v", err)
	}

	r := api.SetupRouter(cfg)
	r.Run(":8080")
}
