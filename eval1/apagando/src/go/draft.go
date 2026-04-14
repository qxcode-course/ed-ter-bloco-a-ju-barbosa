package main
import "fmt"
func main() {

	var id_fila, qtd, sairam int
	fmt.Scan(&id_fila, &qtd, &sairam)
	nafila := make([]int, 0, id_fila)
	
	for i := range nafila {
		fmt.Scan(&nafila[i])
	}

	sairam := make(map[int]int, bool)

	for _, sairam range nafila{
		
	}

	fmt.Printf("%v",id_fila)

}
