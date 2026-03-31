package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Outer x:", x) // prints 10

	{
		x := 20
		fmt.Println("Inner x:", x) // prints 20
	}

	fmt.Println("Outer x after inner scope:", x) // prints 10
}

