package main

import "fmt"

//使用make函数创建切片
func main()  {
demo3()
}
func demo1()  {
	s1:=make([]int,5,10)
	fmt.Println(s1)
	fmt.Println("打印切片长度:",len(s1))
	fmt.Println("打印切片容量:",cap(s1))
	s1[2]=200
	fmt.Println(s1)
	//切片扩容
	s2:=append(s1,3,4,5,6)
	fmt.Println(s2)//如果切片数据超过了规定的容量，那么切片可以自动扩容
	//追加数据
	s3:=append(s1,s2...)
	fmt.Println(s3)
	//遍历切片数据
	for i := 0; i <len(s3) ; i++ {
		fmt.Println(s3[i])
	}
	for a,v:=range s1{
		fmt.Println(a,v) //a index
	}
}
func demo2()  {
	//定义一个切片
	arr:=[]int{0,1,2,3,4,5,6,7,8,9}
	//通过数组创造切片
	s1:=arr[0:8] //start 包含   end 不包含
	s2:=arr[:3]
	s3:=arr[3:]
	s4:=arr[:]
    fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
//数组是值类型  切片是引用类型




}
func demo3()  {
	//深拷贝  值类型  string int float bool array sturct
	//浅拷贝  引用类型  切片、map、
	s1:=[]int{1,2,3,4}
	s2:=make([]int ,0,0)
	fmt.Println(s1,s2)
	for i := 0; i < len(s1); i++ {
	     s2=append(s2,s1[i])
	}
	fmt.Println(s1,s2)
	s2[2]=300
	fmt.Println(s1,s2)
}