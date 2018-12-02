package main

import "fmt"

var channel chan int = make(chan int)
var channel2 chan int = make(chan int)

/*
// deadlock 1
func main() {
	<-channel
}
*/

/*
// deadlock 2
func main() {
	channel<-2
	fmt.Printf("%d\n", 123)
}
*/

func say(para int) {
	fmt.Printf("%d\n", para)
	channel <- <- channel2
}

func say2(para int) {
	// fmt.Printf("%d\n", para)
	channel <- <- channel2
	channel2 <- para
}

func say3(para int) {
	channel <- para
}

func main() {


	/*
	go say(1)
	// 只有这行代码则死锁
	channel <- 3
	*/

	/*
	go say(1)
	// 只有这行代码则不死锁
	channel2 <- 2
	*/

	/*
	// 死锁
	go say2(4)
	fmt.Printf("%d\n", <-channel)
	*/

	go say(1)
	channel2 <-2
	fmt.Printf("%d\n", <-channel)

	/*
	// 不死锁
	go say3(5)
	fmt.Printf("%d\n", <-channel)
	*/

	fmt.Println(len(channel))
	fmt.Println(len(channel2))
	//<- channel
}

