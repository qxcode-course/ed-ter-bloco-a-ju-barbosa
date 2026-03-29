package main
import "fmt"
func main() {
    var fila_inicial, fila_sairam int

    fmt.Scan(&fila_inicial)
    id_inicial := make([]int, fila_inicial) 
    for i := 0; i < fila_inicial; i++ {
        fmt.Scan(&id_inicial[i])
    }

    fmt.Scan(&fila_sairam)
    id_sairam := make([]int, fila_sairam)
    for i := range id_sairam {
        fmt.Scan(&id_sairam[i])
    }

    saiu := make(map[int]bool)

    for _, sairam := range id_sairam {
          saiu[sairam] = true
        
    }

    fila_final := make([]int, 0)

    for _, sairam := range id_inicial {
        if !saiu[sairam] {
            fila_final = append(fila_final, sairam)
        }
    }

    saida := fmt.Sprintf("%v", fila_final)
    fmt.Printf(saida[1 : len(saida)-1])
    fmt.Println(" ")
}
