// If you want a sorted map,
// copy the keys (or values) to a slice,
// sort the slice (using the sort package,(7.6.6)
// and then print out the keys and/or values using a for-range on the slice.
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted: ")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	// a slice of keys is created
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		// appendng that slice of keys
		keys[i] = k
		i++
	}
	// sorting the slice of keys
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("Sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v /", k, barVal[k])
	}

}

// unsorted:
// Key: echo, Value: 56 / Key: golf, Value: 34 / Key: hotel, Value: 16 / Key: juliet, Value: 65 / Key: kilo, Value: 43 / Key: alpha, Value: 34 / Key: bravo, Value: 56 / Key: delta, Value: 87 / Key: lima, Value: 98 / Key: charlie, Value: 23 / Key: foxtrot, Value: 12 / Key: indio, Value: 87 /

// Sorted:
// Key: alpha, Value: 34 /Key: bravo, Value: 56 /Key: charlie, Value: 23 /Key: delta, Value: 87 /Key: echo, Value: 56 /Key: foxtrot, Value: 12 /Key: golf, Value: 34 /Key: hotel, Value: 16 /Key: indio, Value: 87 /Key: juliet, Value: 65 /Key: kilo, Value: 43 /Key: lima, Value: 98 /

// But if what you want is a sorted list you better use a slice of structs, which is more efficient:
