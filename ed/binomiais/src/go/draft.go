package main
import "fmt"

func binomiais(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    return binomiais(n - 1, k - 1) + binomiais(n - 1, k)
}

func main() {
   var n, k int
   fmt.Scan(&n, &k)
   fmt.Println(binomiais(n, k))

}