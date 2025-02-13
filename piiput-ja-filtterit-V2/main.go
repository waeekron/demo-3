package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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

func compute(input []float64, operation string) float64 {
	fmt.Println("in compute()", operation)
	fmt.Println("in compute()", input)
	supportedOperations := operations()
	op, exists := supportedOperations[operation]
	if !exists {
		panic("operation does not exist:/")
	}
	res := op(input...)

	fmt.Printf("Result of numbers %v when (%v)-with each other is equal to %v\n", input, operation, res)
	return res
}

// !!!KÄYTETTY TUKIÄLYÄ!!!

// connects "filters", "pipes them together"
func pipeV2(filters []any, initialData []any) (output any) {
	var wantedArgs []reflect.Value
	var results []reflect.Value
	for i, f := range filters {
		if i == 0 {
			// check what kind of values functions wants as args
			funcType := reflect.TypeOf(f)
			numIn := funcType.NumIn()
			wantedArgs = make([]reflect.Value, numIn)
			// populate slice with zero values
			for j := 0; j < numIn; j++ {
				wantedArgs[i] = reflect.Zero(funcType.In(i))
			}

			// check that initialData can be used as next one's input, panic if next function's parameter list doesn't match with last function's return values
			for k, val := range initialData {
				if reflect.TypeOf(wantedArgs[k]) != reflect.TypeOf(val) {
					panic("Piping error piip-puup")
				}
				wantedArgs[i] = reflect.ValueOf(val)

			}
		} else {
			// check what kind of values functions wants as args
			funcType := reflect.TypeOf(f)
			numIn := funcType.NumIn()
			wantedArgs = make([]reflect.Value, numIn)

			// populate slice with zero values
			for j := 0; j < numIn; j++ {
				wantedArgs[j] = reflect.Zero(funcType.In(j))
			}
			// check that previous function's output can be used as next one's input, panic if next function's parameter list doesn't match with last function's return values
			for k, val := range results {
				if reflect.TypeOf(wantedArgs[k]) != reflect.TypeOf(val) {
					panic("Piping error piip-puup")
				}
				wantedArgs[k] = val

			}
		}
		// call next function with checked argument list
		returnValue := reflect.ValueOf(f).Call(results)
		results = append(results, returnValue...)
		output = results
	}
	return
}

func main() {
	filters := []any{
		readNumbersAndOperation,
		compute,
	}
	pipeV2(filters, []any{})
}
