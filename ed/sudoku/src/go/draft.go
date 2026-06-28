package main
import "fmt"

func resolver(matriz [][]rune, index int) bool {
   N := len(matriz)

   if index == N*N {
    return true
   }

   l := index / N
   c := index % N

   if matriz[l][c] != '.' {
    return resolver(matriz, index+1)
   }
   for v := '1'; v <= rune('0' + N); v++{
    if podeColocar(matriz, l, c, v){
        matriz[l][c] = v
        if resolver(matriz, index + 1){
             return true
        }
        matriz[l][c] = '.'
    }
   }
       return false
    }

func quadrante(matriz [][]rune, l, c int, v rune) bool {
    N:=len(matriz)
    tamanho := 2
    if N == 9 {
        tamanho = 3
    }

    l0 := (l / tamanho) * tamanho
    c0 := (c / tamanho) * tamanho

    for i := 0; i < tamanho; i++{
        for j := 0; j < tamanho; j++{
            if matriz[l0 + i][c0 + j] == v {
                return true
            }
        }
    }
    return false
}

func linha(matriz [][]rune, l int, v rune) bool {
    for _, x := range matriz[l] {
        if x == v {
            return true
        }
    }
    return false
}

func coluna(matriz [][]rune, c int, v rune) bool {
    for _, linha := range matriz {
        if linha[c] == v {
            return true
        }
    }
    return false
}

func podeColocar(matriz [][]rune, l, c int, v rune) bool{
    return !linha(matriz, l, v) &&
           !coluna(matriz, c, v) &&
           !quadrante(matriz, l, c, v)
}

func main() {
   var N int
   fmt.Scan(&N)

   matriz := make([][]rune, N)
   for i := range matriz {
    var linha string
    fmt.Scan(&linha)
    matriz[i] = []rune(linha)
   }

   if resolver(matriz, 0){
    for _, linha := range matriz {
        fmt.Println(string(linha))
    }
   }
}