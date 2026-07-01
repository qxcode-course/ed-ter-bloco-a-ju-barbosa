package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
    value int
    next *Node
    prev *Node
}

type LList struct {
    root *Node
    size int
}

func NewList() *LList {
    dlist := &LList{}
    dlist.root = &Node{}
    dlist.root.next = dlist.root
    dlist.root.prev = dlist.root
    return dlist
}

func (l* LList) Size() int {
    return l.size
}

func (l * LList) String() string{
    if l.size == 0 {
        return "[]"
    }
    var resultado strings.Builder
    n := l.root.next
    fmt.Fprintf(&resultado, "%d", n.value)
    n = n.next

    for i := 1; i < l.size; i++{
        fmt.Fprintf(&resultado, ", %d", n.value)
        n = n.next
    }
    return "[" + resultado.String() + "]"
}

func (l *LList) PushFront(value int){
    n := &Node{
        value: value,
    }  

    primeiro := l.root.next
    n.next = primeiro
    n.prev = l.root
    primeiro.prev = n
    l.root.next = n

    l.size++
}

func (l *LList) PushBack(value int){
    n := &Node{
        value: value,
    }  

    final := l.root.prev
    n.next = l.root
    n.prev = final
    final.next = n
    l.root.prev = n

    l.size++
}

func (l * LList) PopFront(){
    if l.size == 0{
        return
    }
    primeiro := l.root.next
    proximo := primeiro.next
    l.root.next = proximo
    proximo.prev = l.root
    l.size--
}

func (l * LList) PopBack(){
    if l.size == 0 {
        return
    }
    final := l.root.prev
    Final := final.prev
    Final.next = l.root
    l.root.prev = Final
    l.size--
}

func (l *LList) Clear(){
    l.root.next = l.root
    l.root.prev = l.root
    l.size = 0
}
func main() {
scanner := bufio.NewScanner(os.Stdin)
ll := NewList()
    for {
        fmt.Print("$")
        if !scanner.Scan(){
            break
        }
        linha := scanner.Text()
        fmt.Println(linha)
        args := strings.Fields(linha)

        if len(args) == 0 {
            continue
        }

        cmd := args[0]
        switch cmd {
        case "show":
            fmt.Println(ll.String()) 
        case "size":
            fmt.Println(ll.Size())
        case "push_back":
            for _, v := range args[1:] {
                num, _ := strconv.Atoi(v)
                ll.PushBack(num)
            }
        case "push_front":
            for _, v := range args[1:]{
                 num, _ := strconv.Atoi(v)
                ll.PushFront(num)
            }
        case "pop_back":
            ll.PopBack()
        case "pop_front":
            ll.PopFront()
        case "clear":
            ll.Clear()
        case "end":
            return
        default: 
        fmt.Println("fail: comando invalido")
        }
    }

}
