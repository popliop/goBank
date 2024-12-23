package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/popliop/gobank/pkg/database"
	"github.com/popliop/gobank/pkg/server"
)

func main() {

	store, err := database.NewPostgressStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal("Storage init failed", err)
	}

	server := server.NewAPIServer(os.Getenv("SERVER_PORT"), store)
	server.Run()

}

func init() {
	fmt.Println("Initalizing...")
}
