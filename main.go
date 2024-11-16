package main

import "fmt"

func main() {

	person := Prisoner{
		id:     1,
		hasNum: false,
	}
	DumbCheck(person, Maze())

	fmt.Println("Hello, World!")
}
