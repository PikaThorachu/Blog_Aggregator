package main

import (
	"fmt"
	"log"

	"internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %w", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("paul")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
