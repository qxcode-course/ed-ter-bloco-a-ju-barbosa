package main
import "fmt"

type Editor struct {
    texto []rune
    cursor int
    Z [] [] rune
    Y [] [] rune
}

func NewEditor() *Editor{
    return &Editor{texto : []rune{}, cursor: 0}
}

func (e * Editor) salvaEstado(){
    estado := make([]rune, len(e.texto))
    copy(estado, e.texto)
    e.Z = append(e.Z, estado)
    e.Y = nil
}

func (e *Editor) insert(caractareAtual rune){
    e.salvaEstado()
    e.texto = append(e.texto[:e.cursor]), append([]rune{caractereAtual})
}

func main() {
    fmt.Print()

}