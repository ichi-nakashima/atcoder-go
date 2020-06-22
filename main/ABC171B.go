package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	prices := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&prices[i])
	}

	sort.Ints(prices)
	r := 0
	for i := 0; i < K; i++ {
		r += prices[i]
	}

	fmt.Println(r)
}
