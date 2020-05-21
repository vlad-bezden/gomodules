package main

import (
	"fmt"
	su "modules/stringutils"
)

// const state = "New Jersey"

// var name string

// func main() {
// 	name = "Vlad"
// 	from := "USA"
// 	var n int

// 	fmt.Printf("Hello, fellow %s Gophers!\n", state)
// 	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
// 	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
// 	fmt.Printf("By the time %d o'clock comes around, we'll know how to code in Go\n\n", n)
// 	fmt.Println("Let's get started")
// }

func main() {
	fmt.Println("Upper:", su.Upper("upper test"))
	fmt.Println("Lower:", su.Lower("LOWER TEST"))
	fmt.Println("Twice:", su.Twice("Hello"))
}
