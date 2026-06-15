package main
import "fmt"

func NRainhas(n int) int{
    colunas := make(map[int]bool)
    diagonal := make(map[int]bool) //linha - coluna
    diagonal2 := make(map[int]bool) // linha mais coluna
    solucao := 0

    var backtrack func(linha int)
    backtrack = func(linha int){
        if linha == n {
            solucao++
            return
        }

        for coluna := 0; coluna < n; coluna++{
            if colunas[coluna] || diagonal[linha - coluna] || diagonal2[linha + coluna]{
                continue
            }
            colunas[coluna] = true
            diagonal[linha - coluna] = true 
            diagonal2[linha + coluna] = true

            backtrack(linha + 1)
            delete(colunas, coluna)
            delete(diagonal, linha - coluna)
            delete(diagonal2, linha + coluna)
        }
    }
    backtrack(0)
    return solucao
}
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(NRainhas(n))
}