package main

import (
	"fmt"
	"math"
)

//下記のように、インターフェースの値は、
//値と具体的な型のタプルのように考えることができます:
//(value, type)
//インターフェースの値は、特定の基底になる具体的な型の値を保持します。
//インターフェースの値のメソッドを呼び出すと、
//その基底型の同じ名前のメソッドが実行されます。
type A interface {
	B()
}

// struct 構造体
type X struct {
	Y string
}

func (x *X) B() {
	fmt.Println(x.Y)
}

// メソッド
type F float64

func (f F) B() {
	fmt.Println(f)
}

func main() {
	var i A

	i = &X{"hello"}
	describe(i)
	i.B()

	i = F(math.Pi)
	describe(i)
	i.B()
}

func describe(i A) {
	// (value, type) を出力
	fmt.Printf("(%v, %T)\n", i, i)
}
