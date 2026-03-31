package main

import "fmt"

func main() {
	x := 2
	if (x > 3) {
		return
		fmt.Println("x is greater than 3")
	}

	for i := 0; i < 3; i++ {
		continue
		fmt.Println(i)
	}
}
