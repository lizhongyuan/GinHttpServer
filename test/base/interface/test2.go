package main

import "fmt"

type Philosopher interface {
	ThinkingAboutLife()		string
}

type Optimist struct {
	name				string
}

func (optimist Optimist) ThinkingAboutLife() string {
	return "Don't worry, be happy."
}

type Pessimist struct {
	name				string
}

func (pessimist Pessimist) ThinkingAboutLife() string {
	return "I'm a hedonist"
}

func showThemselves(philosopher Philosopher) {
	var method string = philosopher.ThinkingAboutLife()

	fmt.Printf("%s\n", method)
}


func main() {
	var lzy Optimist = Optimist{"lizhongyuan"}
	var yy	Pessimist = Pessimist{"yy"}
	showThemselves(lzy)
	showThemselves(yy)
}