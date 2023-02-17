package main

import "fmt"

func main()  {
/*	下表列出了所有Go语言的逻辑运算符。假定A值为 True，B 值为 False。
	描述
	&&逻辑AND 运算符。如果两边的操作数都是 True，则条件 True，否则为 False     (A&&B)为 False
	||逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。  (A||B)为True
	!逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。   !(A&&B)为True*/
    a:=true
    b:=false
	if a==true&&b==false {
		fmt.Println("心情不好")
	}else{
		fmt.Println("心情还可以")
	}
	if a==true||b==false {
		fmt.Println("心情不好")
	}else{
		fmt.Println("心情还可以")
	}
	if a!=true||b!=false {
		fmt.Println("心情不好")
	}else{
		fmt.Println("心情还可以")
	}
}