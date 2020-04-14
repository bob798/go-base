package main

import (
	"fmt"
	"math"
)

const c string  = "constant"

func main() {

	fmt.Println(c)

	const n  = 5000000000

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	


}
