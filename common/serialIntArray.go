package common

import "strconv"

func SerialIntArray(array []int) (serial string) {
	if len(array) == 0 { return "[]" }

	for i, element := range array {
		if i == 0 {
			serial = "[" + strconv.Itoa(element)
		} else {
			serial = serial + ", " + strconv.Itoa(element)
		}
	}

	return serial + "]"
}