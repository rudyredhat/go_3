// Here is a another more concrete example: we have 2 types stockPosition and car,both have a method getValue();
// realizing that we can define an interface valuable with this method. And the we can define methods that take a
// parameter of type valuable and which are usable by all types that implement this interface, like showValue():
package main

import (
	"fmt"
)

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

// method to determine the value of stockposition
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

// method to determine the value of the car
func (c car) getValue() float32 {
	return c.price
}

// contract that defines diff things that have the val
type valuable interface {
	getValue() float32
}

// anything that satisfies the "valuable" inteface is accepted
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}
func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

/*
Value of the asset is 2308.800049
Value of the asset is 66500.000000
*/
