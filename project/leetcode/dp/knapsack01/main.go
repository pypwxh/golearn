package main

import (
	"fmt"
	"golearn/project/leetcode/dp/infra/algo"
)

func main() {
	var w = []int{1, 2, 3}
	var v = []int{6, 10, 12}
	var C = 5
	//fmt.Println(algo.Knapsack01Recursion(w, v, C))
	//fmt.Println(algo.Knapsack01Recursion02(w, v, C))
	//fmt.Println(algo.Knapsack01Dp01(w, v, C))
	//fmt.Println(algo.Knapsack01Dp02(w, v, C))
	fmt.Println(algo.Knapsack01Dp03(w, v, C))
}
