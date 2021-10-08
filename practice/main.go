package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Gimme a number:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Gimme another number:")
	input2, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	num2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	fmt.Println("Gimme an operator:")
	input3, _ := reader.ReadString('\n')
	for input3 != "-" && input3 != "+" && input3 != "*" && input != "/" {
		input3, _ := reader.ReadString('\n')
	}

	// No ternary operators in Go :(
	if err != nil {
		fmt.Println("you seemed to have messed up: %s\ncomputers don't make mistakes.", err)
		return
	}
	if err2 != nil {
		fmt.Println("you seemed to have messed up: %s\ncomputers don't make mistakes.", err2)
		return
	}
	if err3 != nil {
		fmt.Println("you seemed to have messed up: %s\ncomputers don't make mistakes.", err3)
		return
	}

}

// READING FROM STDIN
// func main() {
// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text:")
// input, _ := reader.ReadString('\n')
// fmt.Println("you said", input)

// fmt.Print("Enter text:")
// numInput, _ := reader.ReadString('\n')
// aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
// if err != nil {
// 	fmt.Println(err)
// } else {
// 	fmt.Println("you entered number", aFloat)
// }
// }

// DATES
// func main() {
// 	n := time.Now()
// 	fmt.Println("Dates and times\ntime.Now() == ", n)

// 	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
// 	fmt.Println("Go launched at ", t)
// 	fmt.Println(t.Format(time.ANSIC))
// 	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
// 	fmt.Printf("parsedTime type: %T\n", parsedTime)
// }

// BASIC MATH
// func main() {
// 	i1, i2, i3 := 11, 512, 12
// 	intsum := i1 + i2 + i3
// 	fmt.Println("int sum:", intsum)

// 	f1, f2, f3 := 21.2, 142.4, 52.5
// 	floatsum := f1 + f2 + f3
// 	fmt.Println("float sum:", floatsum)
// 	floatsum = math.Round(floatsum*100) / 100
// 	fmt.Println("rounded float sum:", floatsum)

// 	circleRadius := 15.5
// 	circumference := circleRadius * 2 * math.Pi
// 	fmt.Printf("circumference: %.2f\n", circumference)
// }

// VARIABLE DECLARATION

// const aConst string = "this is explicitly set const"

// func main() {
// 	var aString string = "explicitly declare variables"
// 	fmt.Println(aString)
// 	fmt.Printf("type: %T\n", aString)

// 	var anInt int = 42
// 	fmt.Println("explicitly declare int: ", anInt)

// 	var defaultInt int
// 	fmt.Println("default int: ", defaultInt)

// 	var anotherStr = "another string"
// 	fmt.Println("implicitly declare vars:\n", anotherStr)
// 	fmt.Printf("type: %T\n", anotherStr)

// 	var anotherInt = 53
// 	fmt.Println(anotherInt)

// 	myString := "this only works inside functions"
// 	fmt.Println("declare vars with `:=`\n", myString)

// 	fmt.Println("consts can be declared outside funcs:\n", aConst)
// }
