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
        
        if i == espada && elemento > 0{
            fmt.Print("%d>", elemento)
        }
        
        if i == espada && elemento < 0{
            fmt.Print("%d<", elemento)
        }
        fmt.Print(elemento)

    fmt.Print("")
    }

    fmt.Print("]\n")
}

func procurar_vivo(jogadores[] int, espada int) int{
    n := len(jogadores)
    for {
       espada = (espada + 1) % n
       if espada < 0 {
        espada += n
       }

       if jogadores[espada] != 0{
        return espada
       }

        }
    }

func main() {
    var qtd, espada int
    fmt.Scan(&qtd, &espada)

    jogadores:= make([]int, 0, qtd)

    for i := 1; i <= qtd; i++{
    jogadores= append(jogadores, i)
}
   espada -= 1
    for range qtd -1{
        print_jogadores(jogadores, espada)
        vai_morrer := procurar_vivo(jogadores, espada)
        jogadores[vai_morrer] = 0
        espada = procurar_vivo(jogadores, espada)
        
    }
    print_jogadores(jogadores, espada)
}
