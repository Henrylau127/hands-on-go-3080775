// packages/imports/begin/main.go
package main

// import the fmt package from the standard library
import "fmt"

// import the time package from the standard library
import "time"

func main() {
	// use the fmt package to print the string "Hello Gopher!"
	fmt.Println("Hello Gopher!")
	// use the time package to print the current weekday
	fmt.Println("Today is", time.Now().Weekday())
}
