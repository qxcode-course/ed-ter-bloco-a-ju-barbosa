package main
import "fmt"

//CODIGO DA PRINCESA V1
//TAVA FAZENDO NO PRINCESA V1
func print_jogadores(jogadores[]int, espada int) {
    fmt.Print("[ ")
    for i, elemento := range jogadores{
        if elemento == 0{
            continue
        }
        if i == espada {
            if elemento > 0 {
                 fmt.Printf("%d>", elemento)
            } else {
                  fmt.Printf("<%d", elemento)
            }
        } else {
            fmt.Printf("%d", elemento)
        }
            fmt.Print(" ")
        }
    fmt.Print("]\n")        
}


func procurar_vivo(jogadores[] int, pos int, direcao int) int{
    n := len(jogadores)
    for {
       pos = (pos + direcao + n) % n
       if jogadores[pos] != 0 {
        return pos
       }
    }
}


func main() {
    var qtd, espada, fase int
    fmt.Scan(&qtd, &espada, &fase)

    jogadores:= make([]int, 0, qtd)
    sinal := fase 

    for i := 1; i <= qtd; i++{
    jogadores= append(jogadores, i*sinal)
    sinal *= -1
}
   espada -= 1
   vivos := qtd

    for vivos > 1{
        print_jogadores(jogadores, espada)

        direcao := 1   
        if jogadores[espada] < 0 {
            direcao = -1
        } 
        alvo := procurar_vivo(jogadores, espada, direcao)
        jogadores[alvo] = 0
        vivos--
        espada = procurar_vivo(jogadores, espada, direcao)

    }
    print_jogadores(jogadores, espada)
}
