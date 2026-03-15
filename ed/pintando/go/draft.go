package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	var p float64
	p = (a + b + c) / 2

	var area float64
	area = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("%.2f\n", area)

}
