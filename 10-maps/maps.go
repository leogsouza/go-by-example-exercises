// Maps are Go built-in associative data types (sometimes called hashes or dicts
// in other languages)
package main

import "fmt"

func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. Println all of its keu/value pairs
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1", v1)

	// The builtin len returns the number of key/value pairs when calld on a map
	fmt.Println("len:", len(m))

	// The builtin delete removes a key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	/* The optional second return value when getting a values from a map indicates
	if the key was present in the map. this can be used to disambiguate between
	missing keys with zero values like 0 or "". Here we didn't need the value
	itself, so we ignored it with the blank identifier _. */
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
