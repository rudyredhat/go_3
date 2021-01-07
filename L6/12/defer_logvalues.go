package main

import (
	"io"
	"log"
)

// defer the function - func()
// naming your return arguments comes in defer statement
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
func main() {
	func1("Go")
}

// 2021/01/07 14:47:10 func1("Go") = 7, EOF
