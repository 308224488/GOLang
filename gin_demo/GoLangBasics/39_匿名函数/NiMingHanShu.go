package main

import "fmt"

func main()  {
	f1()
	f2:=f1
	f2()
	f3:= func(){
		fmt.Println("333333")
	}
	f3()
	func()  {
		fmt.Println("66666")
	}()
}
func f1()  {
	fmt.Println("礼遇")
}
