package main

import "fmt"

func main() {

	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, element := range slice1 {

		if element%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
	// slice2 := make([]int, 0, 10)

	// for i := 0; i <= 10; i++ {
	// 	slice2 = append(slice2, i)
	// }

	// fmt.Println("slice2 ", slice2)
}
