package main

import (
	"basis/server"
	"basis/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db.Connect()
	server.Start()
}
