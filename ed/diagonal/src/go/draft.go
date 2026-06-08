package main
import "fmt"


func diagonal(s string, k int) {
  if len(s) == 0 {
    return
  }
  // imprima k caracteres
  for i := 0; i < k; i++{
    fmt.Print(" ")
  }
  // imprima o primeiro caractere de s e pule a linha
  fmt.Println(string(s[0]))
  // faça a chamada recursiva
  diagonal(s[1:], k + 1)
}

func main() {
    var s string
    fmt.Scan(&s)
    diagonal(s, 0)
}