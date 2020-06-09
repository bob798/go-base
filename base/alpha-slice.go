package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println(v)
		s[i] = 2
	}

	/*值传递 or 引用传递 */
	fmt.Println("end s", s)
	doSomething(s)
	fmt.Println(s)
	doAppend(s)
	fmt.Println(s)

}

/*
append 必须处理接收值
*/
func doAppend(s []int) {
	_ = append(s, 1)
	fmt.Println(s)
}

func doSomething(s []int) {
	if len(s) > 0 {
		s[0] = 10
	}
}
