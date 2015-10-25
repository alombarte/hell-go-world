package main

import(
	"fmt"
	"github.com/alombarte/hell-go-world-math"
)

func main() {
	fmt.Printf("Hello, maths.\n")

	f := hellomath.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, f())
	}
}
