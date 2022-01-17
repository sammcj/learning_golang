package main

import "fmt"

func main() {
	listOfIntegers := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, thisInt := range listOfIntegers {

		if thisInt%2 == 0 {
			fmt.Println(thisInt, "is even")
		} else {
			fmt.Println(thisInt, "is odd")
		}
	}
}
