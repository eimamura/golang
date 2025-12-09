package main

import (
	"fmt"
	"hello/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	message := greetings.Hello2("Gladys")
	fmt.Println(message)
}
