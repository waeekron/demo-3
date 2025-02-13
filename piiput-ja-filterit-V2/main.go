package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

	fmt.Printf("Result of numbers %v when %v-with each other is equal to %v\n", input, operation, res)
	return res
}

type calculatorData struct {
	nums      []float64
	operation string
}

// connects "filters", "pipes them together"
// !!!KÄYTETTY TUKIÄLYÄ!!!
// !!!KÄYTETTY TUKIÄLYÄ!!!
// !!!KÄYTETTY TUKIÄLYÄ!!!
func pipeV2(filters []any, initialData []any) (output any) {
	var wantedArgs []reflect.Value
	var results []reflect.Value
	for i, f := range filters {
		println()
		println()
		println()
		if i == 0 {
			// check what kind of values functions wants as args
			funcType := reflect.TypeOf(f)
			numIn := funcType.NumIn()
			wantedArgs = make([]reflect.Value, numIn)
			// populate slice with zero values
			for j := 0; j < numIn; j++ {
				wantedArgs[i] = reflect.Zero(funcType.In(i))
			}

			// check that previous function's output can be used as new input, panic if next function's parameter list doesn't match with last function's return value
			for k, val := range initialData {
				if reflect.TypeOf(wantedArgs[k]) != reflect.TypeOf(val) {
					panic("Piping error piip-puup")
				}
				wantedArgs[i] = reflect.ValueOf(val)

			}
		} else {
			// check what kind of values functions wants as args
			funcType := reflect.TypeOf(f)
			fmt.Println("current results", results[0].Kind(), results[0].Type())
			fmt.Println(funcType, "funcType ")
			fmt.Println(funcType.In(1), "first arg ")
			fmt.Println(reflect.Zero(funcType.In(1)), "zero value of first arg")
			numIn := funcType.NumIn()
			wantedArgs = make([]reflect.Value, numIn)

			println("looping", numIn, "times")
			// populate slice with zero values
			for j := 0; j < numIn; j++ {
				fmt.Printf("%v", funcType.In(j))
				println("above is the value")
				// println(reflect.TypeOf(wanted))
				wantedArgs[j] = reflect.Zero(funcType.In(j))
				fmt.Printf("%v", wantedArgs[j])
				// fmt.Println(reflect.Zero(funcType.In(j)))
			}
			fmt.Println("zero values after population")
			fmt.Println(wantedArgs[0].Type())
			fmt.Println(wantedArgs[1].Type())
			// check that previous function's output can be used as new input, panic if next function's parameter list doesn't match with last function's return value
			fmt.Println("setting arguments")
			for k, val := range results {
				if reflect.TypeOf(wantedArgs[k]) != reflect.TypeOf(val) {
					panic("Piping error piip-puup")
				}
				fmt.Printf("setting argument arvo : %v \n", val.Interface())
				wantedArgs[k] = val

			}
		}
		println("toimii")
		fmt.Println(wantedArgs)
		for a, v := range wantedArgs {
			fmt.Printf("wantedArgs %v %v %v %v\n", a, v.Interface(), v.Kind(), v.Type)
		}
		println()
		println()
		// call function with the variables as arguments
		fmt.Println(reflect.TypeOf(f))

		fmt.Println("before function call")
		returnValue := reflect.ValueOf(f).Call(results)
		fmt.Println("after function call")
		for _, rv := range returnValue {
			fmt.Println("kissa", rv, rv.Type(), rv.Kind())
			results = append(results, rv)
		}
		output = results
	}
	return
}

func main() {
	// var data calculatorData
	// pipe(
	// 	data,
	// 	readNumbers,
	// 	readOperation,
	// 	computeAndPrint,
	// )

	filters := []any{
		readNumbersAndOperation,
		compute,
	}
	pipeV2(filters, []any{})
}
