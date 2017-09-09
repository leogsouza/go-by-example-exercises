// Go has various value types including strings, integers
// floats, booleans, etc. Here a few basic examples
package main

import "fmt"

func main() {

	// Strings which can be adde together with +
	fmt.Println("go" + "lang")

	// Integer and floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean, with boolean operators as you'd expect
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
