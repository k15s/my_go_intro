package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fibonacci Sequence - Enter a number and have the program generate the
// Fibonacci sequence to that number or to the Nth number.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("to run: 'go run fibonacci.go <int>'\n")
		os.Exit(1)
	} else if n == 1 {
		fmt.Println([1]int{0})
	} else {
		rr := FibonacciRecursive(n)
		rs := FibonacciIterative(n)
		fmt.Println(rr)
		fmt.Println(rs)
	}
}

func FibonacciRecursive(n int) []uint64 {
	res := make([]uint64, n)
	for i := 0; i < len(res); i++ {
		res[i] = FibRecGetVal(i)
	}
	return res
}

func FibRecGetVal(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibRecGetVal(n-1) + FibRecGetVal(n-2)
	}
}

func FibonacciIterative(n int) []uint64 {
	res := make([]uint64, n)
	res[0] = 0
	res[1] = 1
	for i := 2; i < len(res); i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res
}
