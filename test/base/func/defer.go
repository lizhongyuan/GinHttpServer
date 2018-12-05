package main

import (
	"os"
	"fmt"
)


func test1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}


func test2() {
	defer fmt.Printf("%d\n", 1)
	defer fmt.Printf("%d\n", 2)
}

func test3() (result int) {
	defer func() {
		result++
	}()
	return 0
}


func main() {
	// test1("/Users/svenlee/workspace/self/github/data-service/main.go")
	// test2()

	fmt.Printf("%d\n", test3())
}
