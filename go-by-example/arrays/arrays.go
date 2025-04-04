package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[1] = 1
	a[4] = 4
	fmt.Println("emp:", a)

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
