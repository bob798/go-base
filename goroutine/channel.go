package main

import "fmt"

func sum(i []int, c chan int) {
	total := 0
	for _, v := range i {
		total += v
	}
	c <- total
}

func main() {

	a := []int{2, 4, 5, -9, 7, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
	/*-2 11 9
	发出者与接受者无法一一对应
	*/
}
