package main

import (
	"fmt"
	"time"
)

func calcAll(rodPrice []int, num int) {
	memo := make([]int, num)
	cutList := make([]int, num)
	startTime := time.Now()
	for i := 0; i < num; i++ {
		q := -1
		for j := 0; j <= i; j++ {
			var x int
			if j == i {
				x = rodPrice[j]
			} else {
				x = rodPrice[j] + memo[i - j - 1]
			}
			if x > q {
				q = x
				cutList[i] = j
				// fmt.Println("setting cutlist[",i,"] to ",j)
			}
		}
		memo[i] = q
	}
	stopTime := time.Now()
	fmt.Println(stopTime.Sub(startTime))
	fmt.Println("memo",memo)
	fmt.Println("cutList",cutList)
}

func main() {
	rodPrice := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	calcAll(rodPrice, 10)
}
