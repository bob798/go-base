package main

import (
	"fmt"
	"time"
)

/*tab 快捷方式，想你所想*/
func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case i:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*编写时错误编写，输出仍正确，猜猜结果。*/
	switch i {
	case i:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*case语句中，使用逗号分隔多个表达式*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	/*if/else 逻辑另一种形式*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
