package main
import "fmt"
func main() {
  var N, total int
  fmt.Scan(&N, &total)

  vetor := make([]int, total)

  repetidas := -1
  faltam := -1
  

  for i := 0; i < total; i++ {
	fmt.Scan(&vetor[i])
  }

   var repete [100] int 

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++{
			if vetor[i] == vetor[j] {
				repete[i] = 
			}
		}
		
	}


  if repetidas == - 1 {
	fmt.Scan("N\n")
  } else {
	fmt.Printf("%v\n", repetidas)
  }

  if faltam == -1 {
	fmt.Scan("S\n")
  } else {
	fmt.Printf("%v\n", faltam)
  }
}
