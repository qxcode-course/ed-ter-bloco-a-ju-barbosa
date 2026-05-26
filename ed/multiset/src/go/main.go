package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet (capacity int) *MultiSet{//retornando o ponteiro da struct
	return &MultiSet{
		data: make([]int, 0, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (v *MultiSet) expand(){
	if v.capacity == 0{
		v.capacity = 1
		v.data = make([]int, v.capacity)
	} else {
		v.capacity *= 2
		newData := make([]int, v.capacity)
		copy(newData, v.data)
		v.data = newData
	}
}

func (v * MultiSet) Insert(value int) {
	if v.size == v.capacity {
		v.expand()
	}

	index := v.size
	
	for i := 0; i < v.size; i++{
		if value < v.data[i]{
			index = i
			break
		}
	}

	v.data = append(v.data, 0)
	for i := v.size; i > index; i--{
		v.data[i] = v.data[i-1]
	}
	
	v.data[index] = value 
	v.size++
}

func (v * MultiSet) String() string {//método da classe Multiset
	return "[" + Join(v.data[0:v.size], ", ") + "]"

}

func (v *MultiSet) Contains(value int) bool{
	left, right := 0, v.size-1
	for left <= right{
		mid := (left + right)/2
		if v.data[mid] == value {
			return true
		} else if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func (v *MultiSet) Erase(value int) error {
	left, right := 0, v.size-1
	index := -1
	for left <= right{
		mid := (left + right)/2
		if v.data[mid] == value {
			index = mid
			break
		} else if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if index == -1{
		return fmt.Errorf("value not found")
	}

	for i := index; i < v.size - 1; i++{
		v.data[i] = v.data[i+1]
	}

	v.size--
	v.data = v.data[:v.size]

	return nil

}

func (v *MultiSet) Count(value int) int{
	left, right := 0, v.size-1
	index := -1
	for left <= right{
		mid := (left + right)/2
		if v.data[mid] == value {
			index = mid
			break
		} else if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if index == -1{
		return 0
	}

	count := 1
	for i := index - 1; i >= 0 && v.data[i] == value; i--{
		count++
	}

	for i := index + 1; i < v.size && v.data[i] == value; i++{
		count++
	}
	return count
}

func (v *MultiSet) Unique() int {
	if v.size == 0{
		return 0
	}
	count := 1
	for i := 1; i < v.size; i++{
		if v.data[i] != v.data[i-1]{
			count++
		}
	}
	return count
}

func (v *MultiSet) Clear() {
	v.data = make([]int, 0, v.capacity)
	v.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
		    value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
	    	 for _, part := range args[1:] {
		 	 value, _ := strconv.Atoi(part)
			 ms.Insert(value)
			 }
		case "show": 
		fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println(err.Error())
			}
		case "contains":
			 value, _ := strconv.Atoi(args[1])
			 fmt.Println(ms.Contains(value))
		case "count":
			 value, _ := strconv.Atoi(args[1])
			 fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
