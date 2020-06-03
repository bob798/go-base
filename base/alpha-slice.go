package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println(v)
		s[i] = 2
	}

	fmt.Println("end s", s)
	fmt.Println(c)
}
