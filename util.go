package main

func UncheckBoxes (boxes []Box) {
	for i := 0; i < len(boxes); i++ {
		boxes[i].hasBeenChecked = false
	}
}