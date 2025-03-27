package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(a)

	var f string = "apple"
	fmt.Println(f)

	var g string
	g = "banana"
	fmt.Println(g)

	const hundred = 100
	fmt.Println(hundred / 5)
}

func main2() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func main3() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
