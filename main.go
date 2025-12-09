package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Hello())
    hello()
    hello2("Eimura")
}

func hello() {
	fmt.Println("Hello, World!")
}

func hello2(name string) string {
	msg := fmt.Sprintf("Hello, %s", name)
	fmt.Println(msg)
	return msg
}