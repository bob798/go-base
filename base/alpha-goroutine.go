package main

import (
	"fmt"
	"time"
)

func main() {
	// 申请证书
	go applyCert()
	// 注册用户
	register()
}

func applyCert() {
	for i := 0; i < 10; i++ {
		time.Sleep(4 * time.Millisecond)
		fmt.Println("apply cert", i)
	}
}

func register() {
	for i := 0; i < 10; i++ {
		time.Sleep(4 * time.Millisecond)
		fmt.Println("register ", i)
	}
}
