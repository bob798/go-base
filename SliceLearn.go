package main

import "fmt"

func main() {





	/*初始化*/

	s:=make([]string,3)
	fmt.Println("emp", s)

	/*赋值，取值，长度*/
	s[0]="a"
	s[1]= "b"
	s[2]= "c"
	fmt.Println("set", s)
	fmt.Println("get:",s[2])

	fmt.Println("len",len(s))

	/*内建追加，s接受返回的新slice值*/
	s=append(s, "d")
	s=append(s,"e", "f")
	fmt.Println("apd:", s)

	/*写着很舒服*/
	/*copy*/
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy",c)

	/*切片操作*/
	l:=	s[2:5]
	fmt.Println("sl1", l)
	l=	s[:5]
	fmt.Println("sl2", l)
	l=s[2:]
	fmt.Println("sl3", l)

	/*初始化并赋值*/
	t:=	[]string{"g", "h", "i"}
	fmt.Println("dcl", t)

	/*可组成多维数组*/
	twoD	:=make([][]int,3)
	for i := 0; i < 3; i++ {
		innerLen :=	i+ 1
		twoD[i] =make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ", twoD)

}
