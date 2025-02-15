package main

import (
	"log"
	"main.go/core"
	"os"
)

func main() { // we write a SMF file into a buffer and read it back
	file, err := os.ReadFile("Mitsuha's Theme.mid")
	if err != nil {
		log.Fatal(err)
	}
	m := core.New()
	m.Parse(file)
}
