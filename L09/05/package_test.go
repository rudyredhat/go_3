package main

import (
	"fmt"

	"./04/pack1"

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}

// If you need one or more external packages in your application, you will first have to install them
// locally on your machine with the go install command