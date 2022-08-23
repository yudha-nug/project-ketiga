package main

import (
	"fmt"
)

func main() {

	nama := [10]string{"nama 1", "nama 2", "nama 3", "nama 4", "nama 5", "nama 6", "nama 7", "nama 8", "nama 9", "nama 10"}

	for index, element := range nama {

		fmt.Println(index)
		fmt.Println(element)

	}

}
