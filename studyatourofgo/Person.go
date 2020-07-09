package main

import "fmt"

// ポインタとアドレスについて
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{ // pはポインタではなく、Person型の値
		Name: "太郎",
		Age:  20,
	}

	fmt.Printf("最初のp :%+v\n", p)

	p2 := p // Person型の値コピーをして、p2に格納しているにすぎない

	// p2の値を変更しても、それがpに反映されるわけではない。
	p2.Name = "二郎"
	p2.Age = 21
	fmt.Printf("p2で太郎を二郎に書き換えたはずのp :%+v\n", p)

	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p //ポインタ(*Person)をp3に格納している

	// pのアドレス（Personへのポインタである*Person型）を持っているので
	//　p3で書き換えを行うと、その変更はpに反映される。これを参照渡しという。
	p3.Name = "三郎"
	p3.Age = 21
	fmt.Printf("p3でp2を書き換えたp:%+v\n", p)
}
