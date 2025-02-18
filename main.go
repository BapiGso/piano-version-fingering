package main

import (
	"fmt"
	"log"
	"main.go/core"
	"os"
)

func main() { // we write a SMF file into a buffer and read it back
	file, err := os.ReadFile("assets/4007468864.mid")
	if err != nil {
		log.Fatal(err)
	}
	m := core.New()
	m.Parse(file)
	//llm, err := m.Export4LLM()
	//if err == nil {
	//	fmt.Println(llm)
	//}
	llm, err := m.Export4LLM()
	if err != nil {
		return
	}
	fmt.Println(llm)
	m.FillFinger()

}
