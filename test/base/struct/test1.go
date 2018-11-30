package main

import "fmt"




type Person struct {
	id			string;
	name		string;
	level		int32;
}

func (person Person) toString() string {
	return fmt.Sprintf("id=%s,name=%s,level=%d", person.id, person.name, person.level)
}

const LIZONGYUAN_NAME string = "lizhongyuan"

func main() {
	var lizhongyuan Person = Person{"001", LIZONGYUAN_NAME, 0}

	fmt.Printf("%s\n", lizhongyuan.toString())
}


/*
func (person Person) toString() string {
	return fmt.Sprintf("")
}
*/
