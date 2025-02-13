package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sum(values ...float64) float64 {
	var result float64
	for _, val := range values {
		result += float64(val)
	}
	return result
}
func subtract(values ...float64) float64 {
	var result float64
	for _, val := range values {
		result -= float64(val)
	}
	return result
}
func divide(values ...float64) float64 {
	result := values[0]
	for _, val := range values[1:] {
		result /= float64(val)
	}
	return result
}
func multiply(values ...float64) float64 {
	result := 1.0
	for _, val := range values {
		result *= float64(val)
	}
	return result
}

func operations() map[string]func(...float64) float64 {
	return map[string]func(...float64) float64{
		"+": sum,
		"-": subtract,
		"/": divide,
		"*": multiply,
	}

}

func readNumbersAndOperation() ([]float64, string) {
	fmt.Println("Give numbers seperated by spaces")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nums []float64
	for _, val := range strings.Split(line, " ") {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic("bad input!!")
		}
		nums = append(nums, num)
	}
	op := readOperation()
	return nums, op
}

func readOperation() string {
	fmt.Println("Anna operaatio [+ - * /] ")
	var s string
	fmt.Scan(&s)
	return s
}

type operation struct {
	nums []float64
	op   string
}

func main() {
	var wg sync.WaitGroup
	inputCh := make(chan operation)
	resultCh := make(chan float64)
	wg.Add(1)
	go func() {
		// defer wg.Done()
		nums, op := readNumbersAndOperation()
		inputCh <- operation{nums, op}
		close(inputCh)
	}()

	go func() {
		o := <-inputCh
		op, exists := operations()[o.op]
		if !exists {
			panic("ei oo olemassa!!")
		}
		resultCh <- op(o.nums...)
		close(resultCh)
	}()
	println("tulos", <-resultCh)
}
