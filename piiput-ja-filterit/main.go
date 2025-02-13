package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pipe(acc calculatorData, filters ...func(input calculatorData) calculatorData) calculatorData {
	for _, f := range filters {
		result := f(acc)
		acc = result
	}
	return acc
}

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

func readNumbers(input calculatorData) calculatorData {
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

	input.nums = nums
	return input
}

func readOperation(input calculatorData) calculatorData {
	fmt.Println("Anna operaatio [+ - * /] ")
	var s string
	fmt.Scan(&s)
	input.operation = s
	return input
}

func computeAndPrint(input calculatorData) calculatorData {

	supportedOperations := operations()
	op, exists := supportedOperations[input.operation]
	if !exists {
		panic("operation does not exist:/")
	}
	res := op(input.nums...)

	fmt.Printf("Result of numbers %v when %v-with each other is equal to %v\n", input.nums, input.operation, res)
	input.nums = append(input.nums, res)
	return input
}

type calculatorData struct {
	nums      []float64
	operation string
}

func main() {
	var data calculatorData
	pipe(
		data,
		readNumbers,
		readOperation,
		computeAndPrint,
	)
}
