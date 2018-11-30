package main

import "fmt"

type Item struct {
	name		string
	level		int
}


var itemPtr		*Item = &Item{name:"a"}
var item = Item{name:"a"}



func main() {
	fmt.Printf("%p\n", itemPtr)
	fmt.Printf("%s\n", itemPtr.name)

	fmt.Printf("%p\n", &item)
	fmt.Printf("%s\n", item.name)
	fmt.Printf("%#v\n", item)
}
