package main

import (
	"fmt"
)

func main() {

	// var s1 []int
	// // s1[0] = 9
	// fmt.Println(s1, cap(s1), len(s1))
	// s2 := []int{}
	// s2 = append(s2, 9)

	// fmt.Println(s2, cap(s2), len(s2))
	// s2 = append(s2, 10)
	// fmt.Println(s2, cap(s2), len(s2))
	// s2 = append(s2, 11)
	// fmt.Println(s2, cap(s2), len(s2))
	// for i := 0; i < 510; i++ {
	// 	s2 = append(s2, i)
	// }
	// fmt.Println(s2, cap(s2), len(s2))

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3[2:5]
	fmt.Println(s3, cap(s3), len(s3), s4, cap(s4), len(s4))

	s4[0] = 6
	fmt.Println(s3, cap(s3), len(s3), s4, cap(s4), len(s4))

	// create new slice and copy elements
	s4 = append(s4, 10)
	fmt.Println(s3, cap(s3), len(s3), s4, cap(s4), len(s4))

	s4[0] = 9
	fmt.Println(s3, cap(s3), len(s3), s4, cap(s4), len(s4))

	s5 := make([]int, len(s4))
	// Just copy elements
	copy(s5, s4)
	s5 = append(s5, 10)
	fmt.Println(s4, cap(s4), s5, cap(s5))

	/*
		arr := [3]int{1, 2, 3}
		slice := []int{1, 2, 3}
			arr{
				starting Addr,
				total elements
			}


			slice{
					starting Addr,
					total elements(capacity)
					len
			}

			arr := [100]int{1, 2, 3}
			fmt.Println(arr)
			slice := make([]int, 100)	//Lazy allocation
	*/

}
