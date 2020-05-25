package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("to run: 'go run primes.go <int>'\n")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("to run: 'go run primes.go <int>'\n")
		os.Exit(1)
	} else {
		fmt.Println(ComputePrimes(n))
	}
	os.Exit(0)
}

func ComputePrimes(n int) []int {
	primes := make([]int, 0)
	i := 2
	for i < n {
		if n%i == 0 && IsPrime(i) {
			primes = append(primes, i)
		}
		i++
	}
	return primes
}

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
