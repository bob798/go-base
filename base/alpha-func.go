package main

import "fmt"

func main() {
	/*函数参数传递方式
	结论，值传递基本数据类型：值得副本
	引用数据类型：引用副本

	*/

	i := 2
	a := []int{1, 24, 3}
	addInt(i)
	addSlice(a)
	fmt.Println(" i =", i)
	fmt.Println(" a=", a)

	/*
		output
		 i = 2
		 a= [1 1 1]
	*/

}

func addInt(i int) {
	i = 1
}

func addSlice(s []int) {
	for i, _ := range s {
		s[i] = 1

	}
}
