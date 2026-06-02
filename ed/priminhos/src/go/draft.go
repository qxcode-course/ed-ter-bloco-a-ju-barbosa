package main
import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if div == x {
		return true 
	}
	if x % div == 0 {
		return false
	}
	return eh_primo(x, div + 1)
}

func vetor_primos(n int) [] int{
    primos := []int {}
    num := 2

    for len(primos) < n {
        if eh_primo(num, 2){
            primos = append(primos, num)
        }
        num++
    }
    return primos
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(vetor_primos(n))
}