// By this we mean switching the values and keys. If the value type of a map is acceptable as a key
// type, and the map values are unique, this can be done easily:
package main

import "fmt"

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("Inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v/", k, v)
	}
}

// Inverted:
// Key: 65, Value: juliet/Key: 43, Value: kilo/Key: 34, Value: golf/Key: 12, Value: foxtrot/Key: 87, Value: indio/Key: 56, Value: bravo/Key: 16, Value: hotel/Key: 23, Value: charlie/Key: 98, Value: lima/
