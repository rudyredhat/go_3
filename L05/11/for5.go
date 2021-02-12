// continue keyword - continue skips the remaining part of the loop, but then continues with the next
// iteration of the loop after checking the condition,
package main

func main() {

	for i := 0; i < 10; i++ {

		if i == 5 {

			continue
		}
		print(i)
		print(" ")

	}
}

//0 1 2 3 4 6 7 8 9
// 5 is skipped in the output
