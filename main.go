package main

import "fmt"

func main() {

	newMaze := MyMaze {
		mazes: CreateMaze(),
	}

	person := Prisoner{
		id:     1,
		hasNum: false,
	}
	DumbCheck(person, newMaze)

	fmt.Println("Hello, World!")
}
