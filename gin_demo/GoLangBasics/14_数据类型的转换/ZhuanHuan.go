package main

import "fmt"

func main()  {
	//Go语言不支持隐式类型的转换，所有的数据类型的转换都是显式的

	//整型是不能转化为布尔类型的
	a:=5.0
	b:=3.1415926
	c:="爱谁谁"
	d:=int(b)
	fmt.Printf("a:%T\n,b:%T\n,c:%T\nd:%T\n",a,b,c,d)
	fmt.Println(a,b,c,d)
}