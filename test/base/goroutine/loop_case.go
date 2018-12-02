package main


import (
	"fmt"
	"time"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	// loop()
	go loop()
	loop()
	time.Sleep(time.Second)

	var messages chan string = make(chan string)

	go func(message string) {
		messages<-message
	}("Ping!")

	fmt.Println(<-messages)
}

