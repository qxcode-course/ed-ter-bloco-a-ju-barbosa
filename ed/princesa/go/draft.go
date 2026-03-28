package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)
    
    fila := make([]int, 1, n)
    mortos := make(map[int]bool)
    

    for i := range fila {
        fmt.Scan(&fila[i])
    }

    for _, pessoa := range fila {
        if mortos[pessoa] {
           
        } else {
            mortos[pessoa] = true
            
        }
        }
    }
    

}
