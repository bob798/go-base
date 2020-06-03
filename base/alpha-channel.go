package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	equal(c1, c2)
	c3 := c1
	equal(c1, c3)

	equal(c1, nil)

	var c4 chan int
	equal(c4, nil)

	var x, y int

	c1 <- y   // 发送
	x = <-c1  // 接收
	close(c1) // 关闭

}

func equal(c1, c2 chan int) {
	if c1 == c2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal ")
	}

}
