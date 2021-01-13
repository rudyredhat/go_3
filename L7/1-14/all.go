// Arrays and Slices
// var identifier [len]type - var arr1 [5]int
// ch -1 Listing for array
package main

import "fmt"

func main() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {

		arr1[i] = i * 2
	}
	for i := 0; i < len(arr1); i++ {

		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}

// Array at index 0 is 0
// Array at index 1 is 2
// Array at index 2 is 4
// Array at index 3 is 6
// Array at index 4 is 8
//-------------------------------------------------------------------------------------------------------------
// So when an array is passed as an argument to a function like in func1(arr1) , a copy of the array is
// made, and func1 cannot change the original array arr1.
// to change the original arr1 - func1(&arr1)
// ch -2 pointer in array - passing a pointer to a array
package main

import "fmt"

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }
func main() {
	var ar [3]int
	f(ar)
	// passes a copy of ar
	fp(&ar) // passes a pointer to ar
}

// [0 0 0]
// &[0 0 0]
//-------------------------------------------------------------------------------------------------------------
// ch - 3
// When the values (or some of them) of the items are known beforehand, a simpler initialization
// exists using the { , , } notation called array literals (or constructors), instead of initializing every item
// in the [ ]= way.
package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	// 	... indicates the compiler has to count the number of items to obtain the length of the
	// array.
	// var arrLazy = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	//var arrKeyValue = []string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
		// 		Only the items with indexes (keys) 3 and 4 get a real value, the others are set to empty
		// strings, as is shown in the output:
		// Person at 0 is
		// Person at 1 is
		// Person at 2 is
		// Person at 3 is Chris
		// Person at 4 is Ron
	}
}
//-------------------------------------------------------------------------------------------------------------
// ch - 4 
// You can take the address of an array literal to get a pointer to a newly created instance
package main

import "fmt"

func fp(a *[3]int) { fmt.Println(a) }
func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}
// &[0 0 0]
// &[1 1 1]
// &[2 4 8]
//-------------------------------------------------------------------------------------------------------------
// ch -5 
package main

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}
//-------------------------------------------------------------------------------------------------------------
// ch - 6 
// Listing 7.6—array_sum.go - 163 - 7.1.4
// passing an array causes a lot of memory to be used
// use slice or pointer
// sum of numbers in an array
//-------------------------------------------------------------------------------------------------------------
// ch - 7
// 7.7—array_slices.go - 166 
// slice and array diff
// len and cap of slice
// grow a slice
//-------------------------------------------------------------------------------------------------------------
// ch - 8
// 7.2.2 Passing a slice to a function - 168
//-------------------------------------------------------------------------------------------------------------
// ch - 9
// Often the underlying array is not yet defined, we can then make the slice together with the array using the make( ) function:
// var slice1 []type = make([]type, len)
// 7.8—make_slice.go - 169
//-------------------------------------------------------------------------------------------------------------
// ch - 10
// Difference between new() and make()
//-------------------------------------------------------------------------------------------------------------
// ch - 11
// Concatenating strings by using a buffer:
// This method is much more memory and CPU-efficient than +=, especially if the number of strings
// to concatenate is large.
//-------------------------------------------------------------------------------------------------------------
// ch - 12
// use of slice and for range
// 7.9—slices_forrange.go: - 173
// use of slice and for range for strings
// 7.10—slices_forrange2.go - 173
// for range with multidimensional 
//-------------------------------------------------------------------------------------------------------------
// ch - 13
// reslicing on an array where slice is changed
// 7.11—reslicing.go - 175
//-------------------------------------------------------------------------------------------------------------
// ch - 14
// copy and append one slice to another
// 7.12—copy_append_slice.go - 177
//-------------------------------------------------------------------------------------------------------------
