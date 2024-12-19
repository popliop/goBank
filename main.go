package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	store, err := NewPostgressStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal("Storage init failed", err)
	}

	server := NewAPIServer(os.Getenv("SERVER_PORT"), store)
	server.Run()

}

func init() {
	fmt.Println("Initalizing...")
}
