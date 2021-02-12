// Making a slice of bytes from a string
// slice of bytes can be made - c:=[]byte(s)
// else use copy-function copy(dst []byte, src string)
// using for range in slice of strings - 7.13—for_string.go:
package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	// i corresponds the index value and c to the corresponding str
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}

// 0:ÿ 2:界
//--------------------------
// Memory representation of a string and a slice
// s := "hello" ptr->hello and len->5 = hello [5]byte
// t := s[2:3] ptr-> l and len->1 = string
// refer page 178
//--------------------------
// changing char in str
// For example, change “hello” to "cello":
// s:=“hello”
// c:=[]byte(s)
// c[0]=’c’
// s2:= string(c)
// to “cello” : // s2 == “cello”
