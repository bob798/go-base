package main

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println("emp:", a)
	/* var c [5]int string  这样定义会如何，为什么*/

	/*赋值，取值*/
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	/*内置函数len返回数组长度，何为内置？*/
	fmt.Println("len:", len(a))

	/*初始化并赋值*/
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	/*构建多维数组*/
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
