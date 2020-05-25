package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// To save time, test only up to sqrt(n), rounded up. Testing a number n by all
// numbers between 2 and n-1 can quickly become prohibitively time-consuming.
// For instance, if we wanted to check whether 103 is a prime number in this
// way, we would have to divide by 3, 4, 5, 6, 7 ... and so on, all the way to
// 102! Luckily, it's not necessary to test by every single possible factor.
// It's actually only necessary to test the factors between 2 and the square
// root of n. If the square root of n isn't a whole number, round up to the
// nearest whole number and test up to this number instead.
func IsPrime(i int) bool {
	root := int(math.Sqrt(float64(i))) // round
	for a := 2; a <= root; a++ {
		if i%a == 0 {
			return false
		}
	}
	return true
}

func NextPrime(i int) int {
	// reset to odd to conveniently avoid checking evens
	if i%2 == 0 {
		i--
	}
	for {
		i += 2
		if IsPrime(i) {
			return i
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter for next prime - enter STOP to end")
	input := "\n"
	i := 2
	fmt.Printf("%d\n\n", i)
	for input == "\n" {
		i = NextPrime(i)
		fmt.Println(i)
		input, _ = reader.ReadString('\n')
	}
}
