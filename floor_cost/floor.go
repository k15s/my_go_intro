package main

import (
	"fmt"
	"os"
	"strconv"
)

func ParseArgs() (int, int, int) {
	if len(os.Args) != 4 {
		fmt.Println("to run: go run floor.go width height cost")
		os.Exit(1)
	}
	width, ew := strconv.Atoi(os.Args[1])
	height, eh := strconv.Atoi(os.Args[2])
	cost, ec := strconv.Atoi(os.Args[3])
	if ew != nil || eh != nil || ec != nil {
		fmt.Println("to run: go run floor.go width height cost")
		os.Exit(1)
	}
	return width, height, cost
}

func CalcCost(width int, height int, cost int) int {
	return width * height * cost
}

func main() {
	width, height, cost := ParseArgs()
	total := CalcCost(width, height, cost)
	fmt.Printf("total: %d\n", total)
}
