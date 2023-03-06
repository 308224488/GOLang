package main

import "fmt"

func main()  {
	cat:=cataddress{catInFo{"钱多多",20},"杭州","中国"}
   cat.eat()
	cat.age()
    cat.love()
}
func (catInFo catInFo)eat()  {
	fmt.Println(catInFo.name,"说:")
}
//子类重写父类的方法
func (cataddress cataddress)eat()  {
	fmt.Println(cataddress.name,"说:爱我~别走~")
}
//子类扩展自己的方法
func (cataddress cataddress)love()  {
	fmt.Println(cataddress.name,"说:我会永远记得你的~")
}
func (cataddress cataddress)age () {
	fmt.Println("我今年",cataddress.catInFo.age,"了~")
}
type catInFo struct {
	name string
	age  int
}
type cataddress struct {
	catInFo
	city    string
	country string
}
