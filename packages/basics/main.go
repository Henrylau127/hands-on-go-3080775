// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Today is ", time.Now().Weekday())
}
