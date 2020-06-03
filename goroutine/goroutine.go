package main

import (
	"fmt"
	"time"
)

func main() {
	go pint1()
	go pint2()
	go pint3()
	time.Sleep(4 * time.Second)
}

func pint1() {
	fmt.Println(1)
}

func pint2() {
	fmt.Println(2)
}

func pint3() {
	fmt.Println(3)
}
