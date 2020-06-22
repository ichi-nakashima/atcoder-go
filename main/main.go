package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	i := 1
	num := N

	flag := true
	for flag {
		b:= math.Pow(26, float64(i))
		if num > int(b) {
			num = num - int(b)
			i++
		} else {
			flag = false
		}
	}
}
