package main

import "fmt"

var quit_channel chan int

func foo2(id int) {
	fmt.Println(id)
	quit_channel <- 0
}

func main() {
	count := 1000000
	quit_channel = make(chan int)

	for i := 0; i < count; i++ {
		go foo2(i)
	}

	for i := 0; i< count; i++ {
		<- quit_channel
	}
}
