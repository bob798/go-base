package main

import "fmt"

/*
优点：
 未使用的变量强制报错
运行快速
*/
func main() {



	var a string = "inital" //
	fmt.Println(a)

	var b, c int =1 ,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // 变量默认初始化为0

	f := "short"
	fmt.Println(f)




}
