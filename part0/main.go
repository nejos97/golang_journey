package main

import "fmt"

func main() {
	const PI = 3.14
	fullName := "Alexa Merril"
	fmt.Printf("This is my full name %s \n", fullName)
	fmt.Println("Hello world!")

	fmt.Printf("This is my age %v \n", 20)

	fmt.Printf("This is my bank account $%v \n", 200000)

	// this is another block when we want to return the value

	message := fmt.Sprintf("This is the message %v", "value")
	fmt.Println(message)

}
