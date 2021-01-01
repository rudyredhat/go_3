// program shows on which os it runs and the path
// varibale decralation is taken in concern with this program
// page 56 - Lisiting 4.5
package main

import (
	"fmt"
	"os"
)
func main()  {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}