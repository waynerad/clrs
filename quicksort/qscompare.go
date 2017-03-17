package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func recurse_qs_internal(a []float64, minIdx int, maxIdx int) {
	if minIdx == maxIdx {
		return // only 1 element to sort -- nothing to do
	}
	if (maxIdx - 1) == minIdx {
		// only 2 elements to sort -- just swap or not
		if a[minIdx] > a[maxIdx] {
			e := a[minIdx]
			a[minIdx] = a[maxIdx]
			a[maxIdx] = e
		}
		return
	}
	pivot := minIdx
	randpoint := int(rand.Int31n(int32(maxIdx-minIdx))) + minIdx
	e := a[minIdx]
	a[minIdx] = a[randpoint]
	a[randpoint] = e

	for i := minIdx + 1; i <= maxIdx; i++ {
		if a[i] < a[pivot] {
			if i == (pivot + 1) {
				// 2-way swap
				e = a[pivot]
				a[pivot] = a[i]
				a[i] = e
				pivot = i
			} else {
				// 3-way swap
				e = a[pivot+1]
				a[pivot+1] = a[pivot]
				a[pivot] = a[i]
				a[i] = e
				pivot = pivot + 1
			}
		}
	}
	if pivot > minIdx {
		recurse_qs_internal(a, minIdx, pivot-1)
	}
	if pivot < maxIdx {
		recurse_qs_internal(a, pivot+1, maxIdx)
	}
}

func quicksort(a []float64) {
	l := len(a)
	if (l == 0) || (l == 1) {
		return
	}
	recurse_qs_internal(a, 0, l-1)
}

func mine() {
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
	quicksort(a)
	stopTime := time.Now()
	fmt.Println(stopTime.Sub(startTime))
}

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
	fmt.Println("Running my algorithm")
	mine()
	fmt.Println("Running their algorithm")
	theirs()
}
