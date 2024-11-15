package main

import (
	"math/rand"
)

type Box struct {
	id             int
	prisonerNum    int
	hasBeenChecked bool
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

func maze() {
	// create an array of 100 boxes with an id and random prisoner number
	boxes := make([]Box, 100)
	usedRandomNums := make([]int, 100)
	for i := 0; i < 100; i++ {
		// generate a random number between 1 and 100
		// check if the number has already been used
		// if it has, generate a new number
		// if it hasn't, add it to the usedRandomNums array
		// and assign it to the prisonerNum of the box
		// assign the id of the box to the id of the box
		// and set hasBeenChecked to false
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
}
