package main

import (
	"log"
	"os"
)

func main() {
	path, _ := os.Executable()
	b, _ := os.ReadFile("service.conf")
	log.Printf("bin path:     %v\n", path)
	log.Printf("service.conf: %v\n", string(b))
}
