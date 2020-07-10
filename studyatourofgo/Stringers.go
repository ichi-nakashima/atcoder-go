package main

import "fmt"

type Person2 struct {
	Name string
	Age int
}

func (p Person2) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person2{"Arthur Dent", 42}
	z := Person2{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(a.String())
	fmt.Println(z)
	fmt.Printf("%T", a)
}
