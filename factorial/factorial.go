package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		// only the factorial.go module was provided
		fmt.Printf("to run: 'go run factorial.go <int>'\n")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("to run: 'go run factorial.go <int>'\n")
		os.Exit(1)
	} else {
		fmt.Println(FactorialRecursive(n))
		fmt.Println(FactorialIterative(n))
	}
	os.Exit(0)
}

func FactorialRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * FactorialRecursive(n-1)
	}
}

func FactorialIterative(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}
