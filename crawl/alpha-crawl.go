package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//	网址列表
	lists := []string{"https://www.baidu.com", "https://www.sina.com.cn"}

	//	批量抓取内容

	c := make(chan string)
	for _, list := range lists {
		//	抓取
		go craw(list, c)
	}

	//	结果汇总
	web := <-c
	fmt.Println(web)

}

func craw(url string, c chan string) {
	response, err := http.Get(url)
	if err == nil {
		var content []byte
		content, _ = ioutil.ReadAll(response.Body)
		c <- string(content)
	}

}
