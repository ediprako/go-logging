package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("error when open file ", err)
	}

	log.SetOutput(file)
	log.Println("Hello World!")
}
