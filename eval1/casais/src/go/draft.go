package main
import "fmt"
func main() {

    animais := make([]int, 0)
    for i := range animais {
        fmt.Scan(&animais[i])
    }

     solteiros := make(map[int]int,int)

     pares := 0
     for_, solteiros := animais {
        if solteiros[-animais] < 0 {
            solteiros[-animais]--
            pares++
        } else {
            solteiros[-animais]++
        }
     }

     fmt.Printf("%v", pares)

}
