// example with a nested loop (for4.go) break exits the innermost loop:
package main

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
			// prints 012345 for every i
			// as soon as j > 5 breaks and prints space
			// again i is incremented
		}
		print(" ")
	}
}

//012345 012345 012345
