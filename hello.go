package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
