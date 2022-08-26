package main

import (
	"log"

	"github.com/lam.mv/new-server/internal/pkg/config"
)

func main() {
	config := config.GetConfig()
	log.Println(config)
}
