// could not import ./struct_pack/structPack (cannot find package "./struct_pack/structPack" in any of
// 	/usr/lib/golang/src/05/structPack (from $GOROOT)
// 	/home/rpratap/go/src/05/structPack (from $GOPATH))

package main

import (
	"fmt"

	"./05/structPack" // there is a problm with the path, was facing in the previous lession
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.
	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}

// Mi1 = 10
// Mf1 = 16.000000
