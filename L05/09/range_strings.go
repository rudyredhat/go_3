// the for range construct
// breaking string into diff replacememt like char, int, bytes , U+FFFD
package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	for pos, char := range str {
		fmt.Printf("CHaracter on position %d is: %c\n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char,
			pos)
	}
	fmt.Println()
	fmt.Println("index	int(rune) rune char bytes")
	for index, rune := range str {
		fmt.Printf("%-2d	%d	%U	'%c'	%X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
