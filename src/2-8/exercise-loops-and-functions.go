package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for true {
		if ((z * z - x) / (2 * z) * (z * z - x) / (2 * z)) < 1e-10 {
			break;
		}
		z -= (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
