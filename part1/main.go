package main

import "fmt"

// In this module we gonna talk about condition

func main() {
	fmt.Println("Condition in Golang")

	if height := 5; height > 6 {
		fmt.Println("You are tall.")
	} else if height > 4 {
		fmt.Println("You are still tall but not much.")
	} else {
		fmt.Println("You are not tall at all.")
	}
	var myCar car = car{
		Make:   "Toyota",
		Model:  "Yaris",
		Height: 20,
		Width:  30,
	}
	fmt.Println(myCar.Make)
	fmt.Println(myCar.wash())

}

func sub(s1 string, s2 string) string {
	return s1 + s2
}

func add(x, y int) int {
	return x + y
}

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

func (c car) wash() string {
	return fmt.Sprintf("Processing car(%s) washing...", c.Make)
}
