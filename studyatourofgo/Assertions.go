package main

import "fmt"

func main() {
	// 空のinterface は任意の型の値を持つことができる
	var i interface{} = "hello"

	s:= i
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
