package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int //slice de inteiros
}

type Iterator struct { //guarda a posição atual durante a execução
	data  []int //dados
	index int //ponto de começo e parada
}

func NewMyList(values []int) *MyList { //funciona como construtor
	return &MyList{ //cria uma lista nova com os valores que ja existem
		data: values}//retorna a nova struct pra modificar e retornar ela depois 
} 
 

func (l *MyList) Iterator() *Iterator {//l * indica q pertence ao tipo mylist e devolve o endereço (recebedor)
	return &Iterator{
		data: l.data, //dado da mylist
		index: -1}//pra começar do 0
}

func (i *Iterator) HasNext() bool {//checa se há outro 
	return i.index < len(i.data)-1 //retorna true se for vdd
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) { //checa se chegou ao final
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1//se n tiver chegado ao final
	return i.data[i.index]//retorna elemento atual
}

type ReverseIterator struct {
	data  []int//dados (números)
	index int//posição atual
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{
		data: l.data,//dados mylist
		index: len(l.data)}//começa do ultimo
}

func (i* ReverseIterator) HasNext() bool{//retorna false quando chegar em 0
	return i.index > 0//enquanto estiver numa posição maior q 0, existe um elemento atrás
}

func (i* ReverseIterator) Next() int {
	i.index -= 1//diminui 1 posição

	return i.data[i.index]
}


type CyclicIterator struct {
	data  []int//dados (números)
	index int//posição atual
}

func (l *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{
		data: l.data,//dados mylist
		index: -1}//do começo
}

func (i* CyclicIterator) HasNext() bool{
	return true//sempre tem um proximo
}

func (i* CyclicIterator) Next() int {
	i.index += 1//sempre cresce

	if i.index == len(i.data){//se chegar no fim, volta pro 0
		i.index = 0
	}
	return i.data[i.index]
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
			 	fmt.Printf("%v ", it.Next())
			 }
			 fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
		    fmt.Print("[ ")
	        it := mylist.CyclicIterator()
		    for range qtd {
			fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		}
	}

}
