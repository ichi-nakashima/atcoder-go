package main

import "fmt"

type User struct {
	Name string
}

func (u *User) setName() {
	fmt.Printf("u address is %p\n", &u)
	u.Name = "gopher"
}

func main() {
	u := new(User)
	fmt.Printf("u address is %p\n", u)
	u.setName()
	fmt.Println(u.Name)
	fmt.Println("hoge")
}

