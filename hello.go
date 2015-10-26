package main

import (
	"fmt"
	"github.com/alombarte/hell-go-world-math"
)

func main() {
	fmt.Printf("Hello, maths.\n")

	f := hellomath.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, f())
	}

	sum := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of arrays %d = %d\n", sum, hellomath.SumArray(sum))
}
