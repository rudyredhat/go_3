// To demonstrate that the value can be any type,
// here is an example of a map which has a func() int as its value
package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 50 },
	}
	fmt.Println(mf)
}

// map[1:0x4b8c40 2:0x4b8c60 3:0x4b8c80]
// the integers are mapped to addresses of functions
//--------------------
// capacity of maps
// make(map[keytype]valuetype, cap)
// e.g.: map2 := make(map[string]float, 100)
//--------------------
// slices as a map values, means a key corresponds to multiple values
// mp1 := make(map[int][]int)
// mp2 := make(map[int]*[]int)
