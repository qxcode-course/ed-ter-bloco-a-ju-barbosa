package main

import "fmt"

func main() {
	var N [50]int

	var n int
	fmt.Scan(&n)

	for i := 0; i < 0; i++ {
		for j := 0; j < 0; j++ {
			if N[i]%2 == 0 {
				N[i] = 0

			}
		}
	}

	fmt.Printf("%v\n")
}
