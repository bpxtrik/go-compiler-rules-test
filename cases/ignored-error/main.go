package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("nonexistent.txt") 
	fmt.Println(file)

	info, _ := file.Stat()
	fmt.Println(info)
}
