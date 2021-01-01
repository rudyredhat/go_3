// check for the simple implementation of switch statement
// take a int var
// enter the case as the integer val and print the statement
package main

import "fmt"

func main() {
	var num1 int = 100
	switch num1 {
	case 98, 99:
		fmt.Println("It’s equal to 98")
	case 100:
		fmt.Println("It’s equal to 100")
	default:
		fmt.Println("It’s not equal to 98 or 100")
	}
}

//It’s equal to 100
