package main

import (
	"math/rand"
)

type Box struct {
	id             int
	prisonerNum    int
	hasBeenChecked bool
}

type Prisoner struct {
	id     int
	hasNum bool
}

type MyMaze struct {
	mazes []Box
}

// contains checks if a slice contains a specific integer
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func CreateMaze() []Box {
	// create an array of 100 boxes with an id and random prisoner number
	boxes := make([]Box, 100)
	usedRandomNums := make([]int, 100)
	for i := 0; i < 100; i++ {

		randNum := rand.Intn(100) + 1
		for {
			if !contains(usedRandomNums, randNum) {
				usedRandomNums = append(usedRandomNums, randNum)
				break
			} else {
				randNum = rand.Intn(100) + 1
			}

			boxes[i] = Box{i, randNum, false}
		}

	}
	return boxes
}
