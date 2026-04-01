package main
import "fmt"
func main() {
   tamanho := 0
   rot := 0

   fmt.Scan(&tamanho, &rot)

   vetor := make([]int, tamanho)

   for i := 0; i < tamanho; i++{
    fmt.Scan(&vetor[i])
   }

   rotacao := rot % tamanho
   if rotacao < 0{
    rotacao += tamanho
   }

   vet_rotacionando := make([]int, 0, tamanho)

   if rot == 0 {
    fmt.Printf("[")
    for _, i := range vetor {
        fmt.Printf(" %v",i)
    }
    fmt.Printf(" ]\n")
   } else {
   for i := tamanho - rotacao; i < tamanho; i++ {
            vet_rotacionando = append(vet_rotacionando, vetor[i])
        }
    for i := 0; i < tamanho - rotacao; i++ {
        vet_rotacionando = append(vet_rotacionando, vetor[i])
    }

  fmt.Printf("[")
    for _, i := range vet_rotacionando {
        fmt.Printf(" %v", i)
    }
    fmt.Printf(" ]\n")
}
}