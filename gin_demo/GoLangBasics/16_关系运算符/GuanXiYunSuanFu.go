package main

import "fmt"

func main()  {
	//关系运算符
	var  a int =2
	var  b int =2
	fmt.Println(a>b)
	fmt.Println(a<b)
	fmt.Println(a>=b)
	fmt.Println(a<=b)
	fmt.Println(a!=b)
	if a>b {
		fmt.Println("满足条件")
	}else{
        fmt.Println("啥也不是")
	}
}