package main

import "fmt"

// ポインタとアドレスについて
func main() {
	name := "太郎"
	fmt.Printf("name :%v\n", name)

	namePoint := &name // 参照渡し、アドレスを変数に渡した

	// アドレスをそのまま表示
	fmt.Printf("namePoint :%v\n", namePoint)
	// アドレスに格納されている値を表示
	fmt.Printf("namePoint :%v\n", *namePoint)

	*namePoint = "二郎"

	// アドレスをそのまま表示
	fmt.Printf("namePoint :%v\n", namePoint)
	// アドレスに格納されている値を表示
	fmt.Printf("namePoint :%v\n", *namePoint)
	// *namePointには、&name(nameの)が格納されているので、nameも書き換わる
	fmt.Printf("namePoint :%v\n", name)

}
