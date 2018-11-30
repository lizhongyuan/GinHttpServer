package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int = *a
	tmp = *a
	*a = *b
	*b = tmp
}


func main() {
	var a	int = 100
	var b	int = 20

	swap(&a, &b)

	fmt.Printf("a: %d\nb: %d\n", a, b)
}
