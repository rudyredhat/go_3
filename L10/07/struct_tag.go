// A field in a struct can, apart from a name and a type, also optionally have a tag: this is a string attached
// to the field, which could be documentation or some other important label.
// for example: reflect.TypeOf() on a variable gives the right type; if this is a struct type, it
// can be indexed by Field, and then the Tag property can be used.
package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {

		refTag(tt, i)
	}
}
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)    //.TypeOf
	ixField := ttType.Field(ix)     //.Field
	fmt.Printf("%v\n", ixField.Tag) //.Tag
}

// An important answer
// The name of the thing
// How much there are
