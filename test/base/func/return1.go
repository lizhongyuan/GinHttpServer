package main

import "fmt"

func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100

	defer func() {
		x = 1
	}()

	return 0, 0
	// return   //101, 102
	//return x, y  //同上
	//return y, x  //102, 101
}


func change2(a, b int) (int, int) {

	x := 0
	defer func() {
		x = 1
	}()

	return 0, 0
}


func main() {
	a := 1
	b := 2

	c, d := change2(a, b)

	fmt.Println(c, d)
}
