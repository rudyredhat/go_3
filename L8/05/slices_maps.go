// we must use make() 2 times, first for the slice, then
// for each of the map-elements of the slice, like in Listing
package main

import "fmt"

func main() {
	// Version A:
	// page 166 for ref of slice and 168
	items := make([]map[int]int, 5)
	for i := range items {
		// ** for every item we are allocating one memory
		items[i] = make(map[int]int, 1)
		// ** then for each item , the corresponding value is 2
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	// Version B: NOT GOOD
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		// item is only a copy of the slice element
		item[1] = 2
		// This item will be lost on the next iteration
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

// Version A: Value of items: [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
// Version B: Value of items: [map[] map[] map[] map[] map[]]
