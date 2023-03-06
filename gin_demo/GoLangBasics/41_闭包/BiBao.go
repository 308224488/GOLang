package main

import "fmt"

func main()  {
/*	上个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量
	并且该外层函数的返回值就是这个内层函。这个内层函教和外层函数的局部变量，
	统称为闭包结构局部变量的生命周期就会发生改变，正常的局部变量会随着函数的调用而创建，随着函数的结束而销毁
	但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还在维续使用*/
	a1:=increment()
	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(a1())
	a2:=increment()
	fmt.Println(a2())
	fmt.Println(a2())
}
func increment() func() int {
	i:=0 //定义一个局部变量
	//自增并返回变量
  fun:=func ()int{//内层函数，没有执行的
       i++
       return  i
	}
	return fun
}