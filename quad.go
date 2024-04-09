package main

import (
	"fmt"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {

		for col := 1; col <= x; col++ {
			fmt.Print("o")
		}

		fmt.Println()
	}
}

func main() {
	QuadA(5, 3)
}
