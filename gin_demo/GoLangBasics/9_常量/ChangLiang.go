package main

import "fmt"

func main()  {
	//const  userName string="李遇"//使用const定义常量  //显示定义 带数据类型
	const  userId =10//隐式定义  默认推导类型
	//定义多个常量
	const userName,userIdx,sex="李遇",1,true
	fmt.Println(userName,userIdx,sex)

}