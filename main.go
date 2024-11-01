package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slowFunction(ch chan<- int) {
	// just wait for a long time
	time.Sleep(time.Second * 5)
	ch <- rand.Intn(5)
}

func main() {
	fmt.Println("This is before calling the slow function")
	ch := make(chan int)
	go slowFunction(ch)
	fmt.Println("This is after I call the slow function")

	// this is where we get the result back
	result := <-ch
	fmt.Printf("Got %d back\n", result)
}
