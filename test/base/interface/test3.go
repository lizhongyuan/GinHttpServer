package main


import "fmt"


func printAll2(values []string) {

	/*
	for _, val := range values {
		fmt.Println(val)
	}
	*/

	for index, val := range values {
		fmt.Println(index)
		fmt.Println(val)
	}

}


func printAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}


func main(){
	names := []string{"stanley", "david", "oscar"}
	// printAll(names)
	printAll2(names)
}

