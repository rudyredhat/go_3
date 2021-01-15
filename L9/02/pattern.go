// the regexp package
// replace the regexp with a string or using a function
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//string to search
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" // pattern to search for in searchIn
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}
	re, _ := regexp.Compile(pat)
	// replace pat with "##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	// using a function
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

}

// Match found!
// John: ##.# William: ##.# Steve: ##.#
// John: 5156.68 William: 9134.46 Steve: 11264.36
