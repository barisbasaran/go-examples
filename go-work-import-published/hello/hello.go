package main

import (
	"example.com/goodbye"
	"fmt"
	"github.com/barisbasaran/go-publish-modules/greetings"
)

func main() {
	fmt.Println(goodbye.Goodbye())
	fmt.Println(greetings.Hello("Baris"))
}
