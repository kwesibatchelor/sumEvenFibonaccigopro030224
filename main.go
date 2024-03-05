package main

import "fmt"

func sumEvenFibonacci(limit int) int {
	sumEven := 0
	fib1, fib2 := 1, 2

	for fib2 <= limit {
		if fib2%2 == 0 {
			sumEven += fib2
		}

		// generate the next Fibonacci number
		fib1, fib2 = fib2, fib1+fib2
	}
	return sumEven
}

func main() {
	result := sumEvenFibonacci(2000000)
	fmt.Println(result)
}
