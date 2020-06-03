package main

import "fmt"

type User struct {
	age  int
	name string
}

/*

2. 深度对比
reflect DeepEqual
https://sanyuesha.com/2017/12/01/go-goroutine-channel-2/
*/

func main() {
	a := 1
	b := 1

	if a == b {
		fmt.Println("a=b")
	}

	u1 := User{
		age:  1,
		name: "bob",
	}

	u2 := User{
		age:  1,
		name: "bob",
	}

	if u1 == u2 {
		fmt.Println("u1 = 2")
	} else {
		fmt.Println("u1 not equal u2")
	}

}
