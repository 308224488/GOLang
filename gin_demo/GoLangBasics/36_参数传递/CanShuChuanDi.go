package main

import "fmt"

func main()  {
  //定义一个数组
	arr:=[4]int{1,2,3,4}
	result(arr)
	fmt.Println(arr)
	s:=[]int{10,20,30,40,50}
	result1(s)
	fmt.Println(s)
}
/*值传递
arr2 的数期是 从arr1 复制来的,所以是不同的空间
修改 arr2 并不会影啊 arr 1
值传递:传递的是数的副本，修改数据，对于原始的数据没有影响
值类型的数据，默认都是值传递: 基础类型、array、struct
定义一-个数组 [个数]类型*/
func result(arr[4] int)  {
   fmt.Println(arr)
   arr[2]=30
   fmt.Println(arr)
}
//引用传递  切片 slice、map、chan
// 切片 []可以扩容的数组
func result1(s1[] int)  {
	fmt.Println(s1)
	s1[2]=300
	fmt.Println(s1)
}