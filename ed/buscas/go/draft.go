package main
import "fmt"
func main() {
   consultas, buscas := 0,0

   fmt.Scan(&consultas)
   Consulta := make([]string, consultas)
   matriz := make(map[string]int)

   for i := 0; i < consultas; i++{
    fmt.Scan(&Consulta[i])
    matriz[Consulta[i]]++
   }

   fmt.Scan(&buscas)
   Busca := make([]string, buscas)
   resultado := make([]int, buscas)

   for i := 0; i < buscas; i++{
    fmt.Scan(&Busca[i])
    resultado[i] = matriz[Busca[i]]
   }

   saida := fmt.Sprintf("%v", resultado)
   fmt.Println(saida[1 : len(saida)-1])
}
