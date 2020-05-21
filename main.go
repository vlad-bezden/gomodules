package main

import (
	"fmt"
	su "modules/stringutils"
)

func main() {
	fmt.Println("Upper:", su.Upper("upper test"))
	fmt.Println("Lower:", su.Lower("LOWER TEST"))
	fmt.Println("Twice:", su.Twice("Hello"))
}
