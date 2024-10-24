package main

import (
	"auth/app"
	"log"
)

func main() {

	r := app.InitializeApp()

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
