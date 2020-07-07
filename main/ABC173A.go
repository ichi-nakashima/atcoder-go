package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N string
	var i int
	fmt.Scan(&N)

	l := len(N)
	if (l > 3) {
		l = l - 3
	} else {
		l = 0
	}
	i, _ = strconv.Atoi(N[l:])
	if (i == 0) {
		i = 1000
	}
	r := 1000 - i
	fmt.Println(r)
}
