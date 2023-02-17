package main

import (
	"fmt"
)

func main()  {
	//:= 自动推导变量数据类型
	name:="李遇"
	age:=1
	id:=1
	fmt.Printf("name:%T,age:%T,id:%T",name,age,id)
}
