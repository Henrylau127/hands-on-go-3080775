// functions/begin/main.go
package main

import "strconv"
import "errors"
import "fmt"

// simple greet function
//
func greet() string{
	return "Hello"
}

// greetWithName returns a greeting with the name
//
func greetWithName(name string) string {
	return "Hello " + name
}

// greetWithNameAndAge returns a greeting with the name and age of the person
//
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, " + name + ". Age: " + strconv.Itoa(age)
	
	return greeting	
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
//
func divide(int1 int, int2 int) (int, error) {
	if (int2 == 0) {
		return 0, errors.New("Divided By Zero")
	}

	return int1 / int2, nil
}

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("Toni"))
	fmt.Println(greetWithNameAndAge("Toni", 69))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
