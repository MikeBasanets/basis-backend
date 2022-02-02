package main

import (
	"basis/clothing"
	"basis/server"
)

func main() {
	clothing.LoadClothing()
	server.Start()
}
