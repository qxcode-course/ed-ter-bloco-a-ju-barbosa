package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data []int
	size int 
	capacity int
}

func NewSet(capacity int) *Set{
	return &Set{
		data: make([]int, 1),
		size: 0,
		capacity: 1,
	}
}

//igual vetbuild
func (s *Set) Reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	nData := make([]int, newCapacity)
	for i := range s.size {
		nData[i] = s.data[i]
	}
	s.data = nData
	s.capacity = newCapacity
}

//igual bettersearch
func (s *Set) BinarySearch(value int) int {
	comeco, fim := 0, s.size - 1

	for comeco <= fim {
		divisor := (comeco + fim)/2

		if s.data[divisor] > value {
			fim = divisor -1
		} else if s.data[divisor] < value {
			comeco = divisor + 1
		} else {
			return divisor
		}
	}
	return - comeco - 1
	
}

func (s *Set) Insert(value int) {
	indice:= s.BinarySearch(value) 
		if indice >= 0{
			return
		}

		posicao := -indice-1
		if s.size == s.capacity{
			s.Reserve(s.capacity * 2)
		}
		for i := s.size; i > posicao; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[posicao] = value
	s.size++

	}

func (s * Set) Contains(value int) bool {
	return s.BinarySearch(value) >= 0
}

//vetbuild
func (s *Set) Erase(value int) bool {
	indice := s.BinarySearch(value)
		if indice < 0{
			return false
		}

		for i := indice; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return true
	}


//igual vetbuild
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (s *Set) Show() string{
	return "[" + Join(s.data[:s.size], ", ") + "]"
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
			 }
		case "show":
			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value){
				fmt.Println("value not found")
			}
		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 if v.Contains(value){
				fmt.Println("true")
			 } else {
				fmt.Println("false")
			 }
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
