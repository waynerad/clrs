package main

import (
	"fmt"
	"time"
)

func maxPrice(memo []int, cutList []int, rodPrice []int, num int) int {
	if num == 1 {
		// length one -- nothing to cut! no cuts
		memo[1] = rodPrice[0]
		cutList[1] = 1
		return rodPrice[0]
	}
	if num == 2 {
		// length two -- decide to cut or not?
		x := 2 * rodPrice[0]
		if x > rodPrice[1] {
			// cutlist append, how do we do it? -> 1
			memo[2] = x
			cutList[2] = 1 // cut at 1
			return x
		}
		memo[2] = rodPrice[1]
		cutList[2] = 2 // no cut
		return rodPrice[1]
	}
	maxSoFar := -1
	cutMaxPt := -1
	for i := 1; i < num; i++ {
		prevVal := memo[i]
		var prevMax1 int
		var prevMax2 int
		if prevVal < 0 {
			prevMax1 = maxPrice(memo, cutList, rodPrice, i)
		} else {
			prevMax1 = prevVal
		}
		prevVal = memo[num-i]
		if prevVal < 0 {
			prevMax2 = maxPrice(memo, cutList, rodPrice, num-i)
		} else {
			prevMax2 = prevVal
		}
		x := prevMax1 + prevMax2
		if x > maxSoFar {
			maxSoFar = x
			cutMaxPt = i
		}
	}
	if maxSoFar > rodPrice[num-1] {
		memo[num] = maxSoFar
		cutList[num] = cutMaxPt
		return maxSoFar
	}
	memo[num] = rodPrice[num-1]
	cutList[num] = num
	return rodPrice[num-1]
}

func main() {
	rodPrice := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	memo := make([]int, 11)
	cutList := make([]int, 11)
	result := make([]int, 11)
	for i := 0; i < 11; i++ {
		memo[i] = -1
		cutList[i] = 0
		result[i] = 0
	}

	startTime := time.Now()
	for i := range rodPrice {
		result[i+ 1] = maxPrice(memo, cutList, rodPrice, i+1)
	}
	stopTime := time.Now()
	fmt.Println(stopTime.Sub(startTime))
	// fmt.Println("i+1", i+1, "maxPrice(i+1)", maxPrice(memo, cutList, rodPrice, i+1))
	fmt.Println("result",result)
	fmt.Println("cutList",cutList)
}
