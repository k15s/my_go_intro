package main

import (
	"fmt"
	"os"
	"strconv"
)

type ChangeUnit struct {
	coin  string
	value int
	num   int
}

func (cu ChangeUnit) CanSpend(change int) bool {
	if cu.value <= change {
		return true
	} else {
		return false
	}
}

func FloatToChange(n float64) int {
	return int(n * 100)
}

func ParseArgs() (float64, float64) {
	// expecting prog name, cost, and money given
	if len(os.Args) != 3 {
		fmt.Println("to run: go run change.go $cost $payment")
		os.Exit(1) // 1 for error
	}
	cost, ce := strconv.ParseFloat(os.Args[1], 64)
	payment, pe := strconv.ParseFloat(os.Args[2], 64)
	if ce != nil || pe != nil {
		fmt.Println("Something went wrong with your cost/payment")
		os.Exit(1)
	} else if payment < cost {
		fmt.Println("You didn't pay enough")
		os.Exit(1)
	}
	return cost, payment
}

func PrintChange(cus []ChangeUnit) {
	for _, c := range cus {
		fmt.Println(c.coin + ": " + strconv.Itoa(c.num))
	}
}

// both cost and payment are broken down into change, can treat as integers
func FindChange(cost int, payment int) {
	// larges value -> smallest value
	cus := []ChangeUnit{
		ChangeUnit{"quarter", 25, 0},
		ChangeUnit{"dime", 10, 0},
		ChangeUnit{"nickel", 5, 0},
		ChangeUnit{"penny", 1, 0},
	}
	change := payment - cost
	fmt.Println("change owed: " + strconv.Itoa(change))
	for change > 0 {
		for i, c := range cus {
			if c.CanSpend(change) {
				change -= c.value
				cus[i].num += 1
				break
			}
		}
	}
	PrintChange(cus)
}

func main() {
	cost, payment := ParseArgs()
	FindChange(FloatToChange(cost), FloatToChange(payment))
}
