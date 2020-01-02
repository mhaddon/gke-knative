package main

import (
	"./ship"
	"log"
)

func main() {
	log.Print("[Ship] Started.")

	ship.GetPersistence()

	log.Fatalf("[Ship][HTTP] %v",  ship.CreateHTTPListener())
}
