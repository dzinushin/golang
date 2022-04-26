package main

import (
	"github.com/sethvargo/go-password/password"
	"log"
)

func main() {
	// Generate a password that is 16 characters long with 6 digits, 0 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(16, 6, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(res)
}
