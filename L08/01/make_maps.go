// maps = unordered collections of pairs of items, key:values
// declaration - var map1 map[keytype]valuetype = var map1 map[string]int
// ** focus on mapLit and mapCreated - 2 ways for the creation of maps
package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	// built-in function - page 142 - used for allocating memory
	// make = make for built-in reference types (slices, maps, channels)
	mapCreated := make(map[string]float32)
	// mapAssigned is a reference to the same value as mapLit, changing mapAssigned also changes
	// mapLit as the program output shows.
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

// Map literal at "one" is: 1
// Map created at "key2" is: 3.141590
// Map assigned at "two" is: 3
// Map literal at "ten" is: 0
