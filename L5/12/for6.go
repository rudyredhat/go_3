// use of labels with break and continue - goto
// The first word ending with a colon : and preceding the code (gofmt puts it on the preceding line)
// is a label, like LABEL1:
package main

import "fmt"

func main() {
LABEL1:
	// define label
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			// same as using break in the if statement
			// but this is kind of use of break label
			if j == 4 {
				continue LABEL1
				// directly jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

// i is: 0, and j is: 0
// i is: 0, and j is: 1
// i is: 0, and j is: 2
// i is: 0, and j is: 3
// i is: 1, and j is: 0
// i is: 1, and j is: 1
// i is: 1, and j is: 2
// i is: 1, and j is: 3
// i is: 2, and j is: 0
// i is: 2, and j is: 1
// i is: 2, and j is: 2
// i is: 2, and j is: 3
// i is: 3, and j is: 0
// i is: 3, and j is: 1
// i is: 3, and j is: 2
// i is: 3, and j is: 3
// i is: 4, and j is: 0
// i is: 4, and j is: 1
// i is: 4, and j is: 2
// i is: 4, and j is: 3
// i is: 5, and j is: 0
// i is: 5, and j is: 1
// i is: 5, and j is: 2
// i is: 5, and j is: 3
