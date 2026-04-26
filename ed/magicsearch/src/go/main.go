package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	_, _ = slice, value
	comeco, fim := 0, len(slice)-1

	for comeco <= fim {
		meio := (comeco + fim)/2

		if slice[meio] == value {
			comeco = meio + 1 //empurra o comeco
		} else if slice[meio] > value {
			fim = meio -1 //fim vai pro comeco
		} else if slice[meio] < value {
			comeco = meio + 1 //comeco vai pro fim
		} else {
			return meio
		}
	}
	if comeco > 0 && slice[comeco-1] == value { //empurra ate depois de encontrar o valor e então diminui 1 pro valor correto
		
		return comeco-1
	} else {
		return comeco
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
