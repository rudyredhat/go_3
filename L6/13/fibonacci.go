// fibonnaci series - 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765,
package main

import "fmt"

func main() {

	result := 0

	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)

	}
}
func fibonacci(n int) (res int) {

	if n <= 1 {

		res = 1
	} else {

		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
