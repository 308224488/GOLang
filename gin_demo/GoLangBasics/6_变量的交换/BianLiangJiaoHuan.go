package main

import "fmt"

func main()  {
	a:="你吃饭了吗？"
	b:="你睡觉了吗？"
	a,b=b,a
	fmt.Println(a,b)
}