package main

import "fmt"

func myarrays() {

	// controlflow()
	// myarrays()
	//

	var A1 [3]int
	fmt.Println(A1)

	A2 := [3]int{3, 5, 7}
	fmt.Println(A2)

	A3 := A2
	A3[1] = 20
	// A3[3] = 20
	fmt.Println(A3)
	fmt.Println(A2)

	A4 := [5]int{1, 2, 3}
	fmt.Println(A4)

	A5 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(A5)

	// Multidimentional array
	coordinates1 := [3][4]int{
		{0, 0, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1},
	}

	for i := 0; i < len(coordinates1); i++ {
		for j := 0; j < len(coordinates1[i]); j++ {
			fmt.Print(coordinates1[i][j])
		}
	}
	fmt.Println("Capacity: ", cap(A5))

}
