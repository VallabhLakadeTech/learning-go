package main

import "fmt"

const (
	NUM1 = 40
	NUM2
	NUM3 = 50
	NUM4 = 22
)

func main() {

	var num int
	numptr := new(int)
	numptr = &num
	num = 100
	fmt.Println("Making use of New:", *numptr)

	c1 := complex(10.2, 4)
	fmt.Println(c1, real(c1), imag(c1))

	fmt.Printf("Type(c1): %T\n", c1)

	fmt.Println("NUM2: ", NUM2)

}
