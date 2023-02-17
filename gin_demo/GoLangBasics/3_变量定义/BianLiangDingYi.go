package main

import "fmt"

func main()  {
	//定义多个变量
	var(
		id    int
		name  string
		age   int
	)
	/*
	var 形式的声明语句往往是用于需要显式指定变量类型地方
	string默认值为空
	当一个变量被声明之后，如果没有显示的给他赋值，系统自动赋予它该类型的零值:
	整型和浮点型变量的默认值为0和0.0.
	字符串变量的默认值为空字符串。
	布尔型变量默认为 false.
	切片、函数、指针变量的默认为 nil*/
	id=1
	name="李遇"
	age=2
	fmt.Println(id,name,age)
}
