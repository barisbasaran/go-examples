package main

import (
	"fmt"
	"github.com/barisbasaran/go-ws3/greetings"
	"github.com/barisbasaran/pubmodule"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
	fmt.Println(pubmodule.Goodbye())
	fmt.Println(greetings.Hello("Baris"))
}
