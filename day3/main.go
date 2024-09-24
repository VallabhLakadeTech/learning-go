package main

import "fmt"

type BaseClass struct {
}

func (base BaseClass) baseFunc() {
	fmt.Println("I am a base class function")
}

type ChildClass struct {
	baseClass BaseClass
}

func main() {

	childClass := ChildClass{}

	// childClass.baseClass.baseFunc()
	childClass.baseFunc()

}
