package main

import (
	"fmt"
)

func main() {
	var N int
	var M int
	var K int
	fmt.Scan(&N)
	fmt.Scan(&M)
	fmt.Scan(&K)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&B[i])
	}
	// todo2


	// todo1
	count := 1
	for {
		var time int

		if len(A) > 0 && A[0] <= B[0] {
			time = A[0]
			A = append(A[:0], A[1:]...)
		} else if len(B) > 0 {
			time = B[0]
			B = append(B[:0], B[1:]...)
		}
		K -= time
		if K < 0 {
			break
		}
		count++
		if len(A) == 0 && len(B) == 0 {
			break
		}
	}
	fmt.Println(count - 1)
}
