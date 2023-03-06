package main

import "fmt"

func main()  {
  cat:=cat{"钱多多",20,1}
  dog:=dog{"爱多多",20,0}
  cat.eat()
  dog.eat()
}
/*方法
某个类别的行为功能，需要指定的接受者调用。一段独立的功能代码，可以直接调用·
一段独立功能的代码，可以直接调用
命名不能冲突*/
// 方法，多了一个接受者，并且只有这个接受者类型才可以调用
// 接受者 可以是这个类型或者这个类型的指针
func (cat cat)eat(){
	fmt.Println("cat.eat")
}

func (dog dog)eat(){
	fmt.Println("dog.dog")
}
type dog struct {
	name  string
	age   int
	sex   int
}
type cat struct {
	name  string
	age   int
	sex   int
}
// 函数，所有人都可以调用
func demo1()  {
	fmt.Println("函数，所有人都可以调用")
}