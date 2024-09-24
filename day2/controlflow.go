package main

import (
	"fmt"
	"math/rand"
)

func controlflow() {
	randomNum := rand.Intn(100)

	if randomNum < 35 {
		fmt.Println("Fail: ", randomNum)
	} else if randomNum >= 35 && randomNum <= 50 {
		fmt.Println("Needs to work hard: ", randomNum)
	} else if randomNum > 50 && randomNum <= 70 {
		fmt.Println("Can do better:", randomNum)
	} else if randomNum > 70 && randomNum <= 80 {
		fmt.Println("Good:", randomNum)
	} else {
		fmt.Println("Excellent:", randomNum)
	}

	if randomNum := 100; randomNum == 100 {
		fmt.Println("Commendable")
	}

	dayOfWeek := rand.Intn(7)

	switch dayOfWeek {
	case 1:
		fmt.Println("Its Sunday")
	case 2:
		fmt.Println("Its Monday")
	case 3:
		fmt.Println("Its Tuesday")
	case 4:
		fmt.Println("Its Wednesday")
	case 5:
		fmt.Println("Its Thursday")
	case 6:
		fmt.Println("Its Friday")
	case 7:
		fmt.Println("Its Saturday")
		// fallthrough
	default:
		fmt.Println("Invalid")
	}

	myNums := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0
	for i := 0; i < 7; i++ {
		sum = sum + myNums[i]
	}
	fmt.Println("Sum of 1-7:", sum)

	sum = 0
	for _, value := range myNums {
		// sum = sum + myNums[key]
		sum = sum + value
	}
	fmt.Println("Sum of 1-7:", sum)

	sum = 0
	i := 0
	for i < len(myNums) {
		sum = sum + myNums[i]
		i++
	}
	fmt.Println("Sum of 1-7:", sum)

	sum = 0
	i = 0
	for i < len(myNums) {
		sum = sum + myNums[i]
		i++
		if i == 3 {
			continue
		}
		if i == 7 {
			break
		}

	}
	fmt.Println("Sum :", sum)
}
