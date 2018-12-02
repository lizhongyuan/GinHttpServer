package main


import (
	"fmt"
)


var messages chan string = make(chan string)

func pushMsg(message string) {
	messages<-message
}


func main() {

	go pushMsg("Ping")

	fmt.Printf("%s\n", <-messages)
}

