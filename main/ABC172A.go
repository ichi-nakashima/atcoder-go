package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	r := a + a*a + a*a*a
	fmt.Println(r)
}
