package main

import (
	"fmt"
	"log"

	"internal/config"
)

func main() {
	username := "paul"

	err := config.SetUser(username)
	if err != nil {
		log.Fatalf("Failed to update: %v", err)
	}

	updatedconfig, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read updated config: %v", err)
	}
	fmt.Println(updatedconfig.DbUrl)
}
