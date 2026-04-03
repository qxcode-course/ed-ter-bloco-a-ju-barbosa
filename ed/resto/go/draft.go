package main
import "fmt"

    func dividir(num int) {
        if (num == 0) {
            return
        }
     divisao := num/2
     resto := num % 2
    dividir(divisao)
    fmt.Printf("%v %v\n", divisao, resto)
    }

    func main(){
        var num int
        fmt.Scan(&num)
        (dividir(num))
    }

