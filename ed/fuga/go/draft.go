package main
import "fmt"

func main() {
   var H, P, F, D int
   fmt.Scan(&H, &P, &F, &D)

   if H > P && F < P && F < H {
      if D == -1 {
         fmt.Printf("S\n")
      } else {
         fmt.Printf("N\n")
      }
   } else if H < P && F > P && F < H {
      if D == 1 {
         fmt.Printf("S\n")
      } else {
         fmt.Printf("N\n")
      }
   } else if H < F && F < P && F > H{
      if D == -1 {
         fmt.Printf("S\n")
      } else {
         fmt.Printf("N\n")
      }
   } else if H > P && F > P && F > H {
      if D == -1 {
         fmt.Printf("S\n")
      } else {
         fmt.Printf("N\n")
      }
   } else if  H > P && F > P && F < H {
      if D == 1 {
         fmt.Printf("S\n")
      } else {
         fmt.Printf("N\n")
      }
   } 
}