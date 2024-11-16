package main

import "fmt"

func UncheckBoxes(boxes []Box) {
	for i := 0; i < len(boxes); i++ {
		boxes[i].hasBeenChecked = false
	}
}

// takes prisoner reference and takes maze object and does a random check and
func DumbCheck(x Prisoner, y MyMaze) {

	for i := 0; i < len(y.mazes); i++ {
		fmt.Println(y.mazes[i])
	}

}
