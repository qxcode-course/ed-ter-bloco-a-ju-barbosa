package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	_ = vet
	if len(vet) == 0 { 
		return "[]"
	}

	if len(vet) == 1 {
		return fmt.Sprintf("[%d]", vet[0])
	}
	return fmt.Sprintf("[%d, %s", vet[0], tostr(vet[1:])[1:])
	}
	

func tostrrev(vet []int) string {
	_ = vet
	if len(vet) == 0 {
        return "[]"
    }

	if len(vet) == 1 {
        return fmt.Sprintf("[%d]", vet[0])
    }

	return fmt.Sprintf("[%d, %s", vet[len(vet)-1], tostrrev(vet[:len(vet)-1])[1:])
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	_ = vet
	if len(vet) <= 1 {
		return
	}

	vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
	reverse(vet[1 : len(vet)-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return 0
	}

	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet

	if len(vet) == 0 {
		return 1
	}

	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet

	if len(vet) == 0 {
        return -1
    }

    // função recursiva interna que retorna índice e valor
    var rec func(v []int, offset int) (int, int)
    rec = func(v []int, offset int) (int, int) {
        if len(v) == 1 {
            return offset, v[0]
        }
        idx, val := rec(v[1:], offset+1)
        if v[0] < val {
            return offset, v[0]
        }
        return idx, val
    }

    idx, _ := rec(vet, 0)
    return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
