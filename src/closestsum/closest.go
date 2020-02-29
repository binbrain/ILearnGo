// 5 Problem Solving Tips for Cracking Coding Interview
// Questions. https://www.youtube.com/watch?v=GBuHSRDGZBY

package main

import (
	"fmt"
)

type pair struct {
	n1 int
	n2 int
}

// O(n^2)
func ClosestSumBruteForce(arr1, arr2 []int, find int) (p pair) {
	closest := -1
	for _, n1 := range arr1 {
		for _, n2 := range arr2 {
			try := (n1 + n2) - find
			if try < 0 {
				try = -try
			}
			if try < closest || closest < 0 {
				closest = try
				p = pair{n1, n2}
			}
		}
	}
	return p
}

// O(x*n)
func ClosestSumBetter(arr1, arr2 []int, find int) (p pair) {
	set := make(map[int]bool)
	for _, n1 := range arr1 {
		set[n1] = true
	}
	for x := 0;;x++ {
		for _, n2 := range arr2 {
			compliment := find - n2
			if _, ok := set[compliment + x]; ok {
				return pair{compliment + x, n2}
			}
			if _, ok := set[compliment - x]; ok {
				return pair{compliment - x, n2}
			}
		}
	}
}

func main() {
    fmt.Println(ClosestSumBruteForce([]int{2, 5, 8}, []int{2, 4, 8}, 8))
    fmt.Println(ClosestSumBetter([]int{2, 5, 8}, []int{2, 4, 8}, 8))
    //fmt.Println(ClosestSumBetter([]int{}, []int{}, 10))
}
