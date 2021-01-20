// The general format of the definition of a struct is as follows:
// type identifier struct {
// 		 field1 type1
// 		 field2 type2
// 		 ...
// }
// Memory for a new struct variable is allocated with the new function, which returns a pointer to the
// allocated storage: var t *T = new(T)
package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	//structname.fieldname = value
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}

// The int is: 10
// The float is: 15.500000
// The string is: Chris
// &{10 15.5 Chris}
