package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("it's int. %v\n",v)
	case string:
		fmt.Printf("it's string. %v\n",v)
	case bool:
		fmt.Printf("it's bool. %v\n",v)
	default:
		fmt.Printf("it's default. %v\n",v)
	}
}
func main() {
	do(12)
	do("hoge")
	do(true)
}
