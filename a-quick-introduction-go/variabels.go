package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var anotherGlobal = -5678

func main() {
	var j int 
	i :=  Global + anotherGlobal
	fmt.Println("initial j value",j)
	j = Global
	// math.Abs() requires a float64 parameter
 	// so we type cast it appropriately
 	k := math.Abs(float64(anotherGlobal))
 	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}