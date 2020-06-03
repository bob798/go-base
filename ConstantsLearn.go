package main

import (
	"fmt"
	"math"
)

/*
变量使用方便，简单，类型无需转换
*/

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
