package main

import (
	"errors"
	"fmt"

	"my_go_intro/sorting_algorithms/sorts/int_merge_sort"
	"my_go_intro/sorting_algorithms/sorts/merge_sort"
)

// in this module's directory '~/go-projects/src/my_go_intro/sorting_algorithms'
// where '~/go-projects' is the GOPATH:
//
// go install
// parent dir name 'sorting_algorithms' should be an executable in 'GOPATH/bin'
//
// OR
//
// go build main.go
// main should be an executable in the current dir
//
// OR
//
// go run main.go

type ReversedInt struct {
	value int
}

// note how this new type fulfills my merge_sort interface
type ReversedIntSlice []ReversedInt

func (r ReversedIntSlice) Get(i int) interface{} { return r[i] }
func (r ReversedIntSlice) Len() int              { return len(r) }
func (r ReversedIntSlice) Less(i, j int) bool    { return r[i].value > r[j].value }
func (r ReversedIntSlice) Set(i int, e interface{}) {
	// need to give j a type assertion
	if e_i, ok := e.(ReversedInt); ok {
		r[i] = e_i
	} else {
		errors.New("Type error in merge_sort")
	}
}

func main() {
	fmt.Println("sorting_algorithms main.go running...")

	i_a := []int{4, 3, 2, 1}
	int_merge_sort.IntMergeSort(i_a)
	fmt.Println(i_a)

	i_b := []int{3, 6, 1, 7, 89, 5, 23, 12, 78, 93, -1, 44, 25, 31}
	int_merge_sort.IntMergeSort(i_b)
	fmt.Println(i_b)

	a := []int{5, 7, 1, 2, 0, -3, 45, 23}
	merge_sort.MergeSort(merge_sort.IntSlice(a))
	fmt.Println(a)

	ris_1 := ReversedIntSlice{ReversedInt{1}, ReversedInt{2}, ReversedInt{3}, ReversedInt{4}}
	merge_sort.MergeSort(ris_1)
	fmt.Println(ris_1)

}
