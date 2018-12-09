package main

import (
	"fmt"
	"math"
)

// func Sqrt(x float64) float64 {
// 	z := 1.0
//
// 	for count := 0; count < 10; count += 1 {
// 		z -= (z*z - x) / (2 * z)
// 	}
//
// 	return z
// }

// Sqrt returns square root of x
func Sqrt(x float64) (float64, int) {
	z0, z := 1.0, x
	count := 0

	for math.Abs(z-z0)/z > 1e-15 {
		z0, z = z, z-(z*z-x)/(2*z)
		count++
	}

	return z, count
}

func main() {
	mathSqrt2 := math.Sqrt(2)
	sqrt2, count := Sqrt(2)

	fmt.Println("math.Sqrt(2)", mathSqrt2)
	fmt.Println("Sqrt(2)", sqrt2)
	fmt.Println("Diff", mathSqrt2-sqrt2)
	fmt.Println("Count", count)
}
