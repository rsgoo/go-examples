package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("example03.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	for i := 0; i < 10; i++ {
		log.Printf("log iteration %d\n", i)
	}
}
