package main

import (
	"bufio"
	"fmt"
	"os"
)


type Pos struct {
	l,c int
}

func (p Pos) getNeigh() []Pos{

}
func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	
	//coloca o elemento a ser analisado 

	//enquanto tem alguem na pilha
	// tira o elemento da pilha, se ele é arvore queima ele. se n for continue

	//topo = o topo da pilha
	//vejo se tem vizinhos
	//se tiver alguem q é arvores e nao foi visitado queimo
	//marca visitado e push vizinho
	//continue


	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
