package merge_sort

import (
	"errors"
)

// A type, typically a collection, that satisfies my custom sort interface can
// be sorted by the routines in this package.  The methods require that the
// elements of the collection be enumerated by an integer index.
type NonInPlaceSortInterface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool

	// Gets the element at index i in the collection
	Get(i int) interface{}

	// Sets the element at index i in the collection with element e
	Set(i int, e interface{})
}

func MergeSort(data NonInPlaceSortInterface) {
	MergeSortHelper(data, 0, data.Len()-1)
}

func MergeSortHelper(data NonInPlaceSortInterface, lo int, hi int) {
	if hi <= lo {
		return
	} else {
		med := lo + (hi-lo)/2
		MergeSortHelper(data, lo, med)
		MergeSortHelper(data, med+1, hi)
		Merge(data, lo, med, hi)
	}

}

func Merge(data NonInPlaceSortInterface, lo int, med int, hi int) {
	// slice: array pointer + length, initialized in memory with length 0
	s := make([]interface{}, 0, hi-lo)
	i := lo
	j := med + 1
	for i <= med || j <= hi {
		if i > med {
			s = append(s, data.Get(j))
			j++
		} else if j > hi {
			s = append(s, data.Get(i))
			i++
		} else if data.Less(i, j) {
			s = append(s, data.Get(i))
			i++
		} else {
			s = append(s, data.Get(j))
			j++
		}
	}
	i = lo
	for _, e := range s {
		data.Set(i, e)
		i++
	}
}

// Now let's bind my custom sort interface to some common use cases

type IntSlice []int

func (p IntSlice) Get(i int) interface{} { return p[i] }
func (p IntSlice) Len() int              { return len(p) }
func (p IntSlice) Less(i, j int) bool    { return p[i] < p[j] }
func (p IntSlice) Set(i int, e interface{}) {
	// need to give j a type assertion
	if e_i, ok := e.(int); ok {
		p[i] = e_i
	} else {
		errors.New("Type error in merge_sort")
	}
}
