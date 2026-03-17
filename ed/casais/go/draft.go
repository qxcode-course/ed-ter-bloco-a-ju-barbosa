package main

import "fmt"

func main() {
	var numero int
	fmt.Scan(&numero)

	vetor := make([]int, numero)

	for i := 0; i < numero; i++ {
		fmt.Scan(&vetor[i])

	}

}
