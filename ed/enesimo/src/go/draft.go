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

func enesimo(x int, atual int, contador int) int{
    if contador == x {
        return atual
    }
    proximo := atual + 1 
    if eh_primo(proximo, 2){
        return enesimo(x, proximo, contador + 1)
    }
    return enesimo(x, proximo, contador)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(enesimo(x , 2, 1))
}
