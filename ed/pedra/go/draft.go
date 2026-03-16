package main

import "fmt"

func main() {
	var N, A, B, diferenca int

	fmt.Scan(&N)
	vetor := make([]int, N)

	valor := -1
	menor := 20000

	for i := 0; i < N; i++ {
		fmt.Scan(&A, &B)

		if A >= 10 && B >= 10 {
			diferenca = A - B
			if diferenca < 0 {
				diferenca *= -1
            vetor[i] = diferenca
			} else {
				vetor[i] = diferenca
			}

			if vetor[i] < menor {
				menor = vetor[i]
				valor = i
			}

		} else {
			vetor[i] = 10000
		}
	}

	if valor == -1 {
		fmt.Printf("sem ganhador\n")
	} else {
		fmt.Printf("%v\n", valor)
	}

}
