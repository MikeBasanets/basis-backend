package main

import (
	"basis/db"
	"basis/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db.Connect()
	server.Start()
}
