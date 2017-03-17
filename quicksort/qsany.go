// This is the 3rd and final program in the series of quicksort programs I wrote.
// The first one (qs.go) was to see if I could implement quicksort under
// "competition conditions" by knowing only the problem (sort numbers) and that
// the quicksort algorithm works by a "divide-and-conquer" methodology -- pick a
// pivot value, rearrange everything so that everthing on one side is above and
// everything on the other side is below the pivot, then recurse on the two
// halves. Everything else I implemented as fast as I could, and I had a working
// program in 42 minutes. That's pretty good, but probably not good enough to
// win Google Code Jam (if quicksort was ever a Code Jam problem, which it can't
// be of course because it's too well known).
//
// The 2nd program, qscompare.go, is just a copy of the first program with a
// function to compare with the built-in sort system in Go. When I found mine
// was faster (by a lot), I wrote the third program. The third program switches
// from a program to sort floating-point numbers to a program to sort "anything"
// by abstracting out the compare and swap operations to functions called through
// an interface. There are two compare and two sort functions because the
// interface abstracts the sort function exactly as I originally wrote it, and I
// did two different comparisons (greater than and less than -- no greater than
// or equal or less than or equal) and two swap functions (one that swaps 2 and
// one that swaps 3 -- I'm sure this is a weird quirk of the way I implemented
// quicksort, but it works).
//
// Once abstracted through the interface, my times are about the same as Go's
// built-in sort system. Actually, Go's times are surprisingly consistent while
// mine vary more, and when mine happens to get a fast time, it slightly beats
// Go's, and when it has a slower time, Go beats it, sometimes by a lot. I'm not
// sure what the Go developers did to get such consistent times but it would
// appear they know what they are doing.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// This is the interface that lets you sort anything
type mySorter interface {
	sizeOfArray() int
	compareGt(minIdx int, maxIdx int) bool
	compareLt(i int, pivot int) bool
	swap2(minIdx int, maxIdx int)
	swap3(pivot int, i int)
}

// This is type and the functions below that use it as a receiver define my object for sorting an array of floats
type mySortObj struct {
	stuff []float64
}

func (objPtr *mySortObj) sizeOfArray() int {
	return len(objPtr.stuff)
}

func (objPtr *mySortObj) compareGt(minIdx, maxIdx int) bool {
	if objPtr.stuff[minIdx] > objPtr.stuff[maxIdx] {
		return true
	}
	return false
}

func (objPtr *mySortObj) compareLt(i int, pivot int) bool {
	if objPtr.stuff[i] < objPtr.stuff[pivot] {
		return true
	}
	return false
}

func (objPtr *mySortObj) swap2(minIdx int, maxIdx int) {
	e := objPtr.stuff[minIdx]
	objPtr.stuff[minIdx] = objPtr.stuff[maxIdx]
	objPtr.stuff[maxIdx] = e
}

func (objPtr *mySortObj) swap3(pivot int, i int) {
	// 3-way swap
	e := objPtr.stuff[pivot+1]
	objPtr.stuff[pivot+1] = objPtr.stuff[pivot]
	objPtr.stuff[pivot] = objPtr.stuff[i]
	objPtr.stuff[i] = e
}

// The actual quicksort implementation
func recurse_qs_internal(objPtr mySorter, minIdx int, maxIdx int) {
	if minIdx == maxIdx {
		return // only 1 element to sort -- nothing to do
	}
	if (maxIdx - 1) == minIdx {
		// only 2 elements to sort -- just swap or not
		if objPtr.compareGt(minIdx, maxIdx) { // true if array[minIdx] > array[maxIdx]
			objPtr.swap2(minIdx, maxIdx)
		}
		return
	}
	pivot := minIdx
	randpoint := int(rand.Int31n(int32(maxIdx-minIdx))) + minIdx
	objPtr.swap2(minIdx, randpoint)

	for i := minIdx + 1; i <= maxIdx; i++ {
		if objPtr.compareLt(i, pivot) { // true if array[i] < array[pivot]
			if i == (pivot + 1) {
				objPtr.swap2(pivot, i)
			} else {
				objPtr.swap3(pivot, i)
				pivot = pivot + 1
			}
		}
	}
	if pivot > minIdx {
		recurse_qs_internal(objPtr, minIdx, pivot-1)
	}
	if pivot < maxIdx {
		recurse_qs_internal(objPtr, pivot+1, maxIdx)
	}
}

// A thunk function to get things started so the user doesn't have to figure out the min and max index parameters
func quicksort(objPtr mySorter) {
	l := objPtr.sizeOfArray()
	if (l == 0) || (l == 1) {
		return
	}
	recurse_qs_internal(objPtr, 0, l-1)
}

// My version
func mine() {
	currentTime := time.Now()
	currentUnix := currentTime.UnixNano()
	rand.Seed(currentUnix)

	size := 8000000
	fmt.Println("size is ", size)
	fSize := float64(size)

	var objPtr mySortObj
	objPtr.stuff = make([]float64, size)
	for i := 0; i < size; i++ {
		objPtr.stuff[i] = rand.Float64() * fSize
	}
	startTime := time.Now()
	quicksort(&objPtr)
	stopTime := time.Now()
	fmt.Println(stopTime.Sub(startTime))
}

// Calling Go's sort system through the "sort" package for time comparison
func theirs() {
	currentTime := time.Now()
	currentUnix := currentTime.UnixNano()
	rand.Seed(currentUnix)

	size := 8000000
	fmt.Println("size is ", size)
	fSize := float64(size)

	var a []float64
	a = make([]float64, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Float64() * fSize
	}
	startTime := time.Now()

	sort.Float64s(a)

	stopTime := time.Now()
	fmt.Println(stopTime.Sub(startTime))
}

func main() {
	fmt.Println("Running mine:")
	mine()
	fmt.Println("Running theirs:")
	theirs()
}
