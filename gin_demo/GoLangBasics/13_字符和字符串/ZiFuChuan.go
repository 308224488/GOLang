package main

import "fmt"

func main()  {
	test_01:='多'
	test_02:="我的钱多多啊，我真舍不得你！"
	fmt.Printf("test_01:%T,test_02:%d\n",test_01,test_01)
	fmt.Printf("test_01:%T,test_02%s",test_02,test_02)
	//字符串连接
	fmt.Printf("你啊"+",到底想要什么？")
	//转义字符
	fmt.Printf("我你/“ 🐎")
}