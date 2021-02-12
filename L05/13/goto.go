// use of break label -to break out of a switch
// goto keyword can be used to follow up by the label name

package main

func main() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

//01234
