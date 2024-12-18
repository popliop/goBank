package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	store, err := NewPostgressStorage()
	if err != err {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(os.Getenv("SERVER_PORT"), store)
	server.Run()

}

//https://whatismyipaddress.com/ip/83.191.174.253
