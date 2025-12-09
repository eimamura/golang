package main

import (
	"fmt"
	"hello/greetings"
	"rsc.io/quote"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println(quote.Hello())
	message := greetings.Hello2("Gladys")
	fmt.Println(message)

	wg.Wait()
}
