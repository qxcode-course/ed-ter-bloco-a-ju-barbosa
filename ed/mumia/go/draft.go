package main
import "fmt"
func main() {
   var a int
   var nome string
   
   fmt.Scan(&nome)
   fmt.Scan(&a)

   if a < 12 {
    fmt.Printf("%v eh crianca\n", nome)
   } else if  a < 18{
    fmt.Printf("%v eh jovem\n", nome)
   } else if a < 65 {
    fmt.Printf("%v eh adulto\n", nome)
   } else if  a < 1000 {
    fmt.Printf("%v eh idoso\n", nome)
   } else {
    fmt.Printf("%v eh mumia\n", nome)
   }
}
