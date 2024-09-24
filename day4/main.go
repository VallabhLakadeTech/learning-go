package main

import "fmt"

type House struct {
	door       int
	galary     int
	windows    int
	rooms      int
	carpetArea int
}

func main() {

	house1 := House{
		door:       2,
		galary:     2,
		windows:    5,
		rooms:      3,
		carpetArea: 1025,
	}

	fmt.Println(house1)

}

func (house House) String() string {
	return fmt.Sprintf("This house contains %v door, %v galaries, %v windows and %v rooms with a carpet area of %v", house.door, house.galary, house.windows, house.rooms, house.carpetArea)
}
