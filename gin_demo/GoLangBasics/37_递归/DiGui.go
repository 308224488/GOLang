package main

import "fmt"

func main()  {
  s:= getSum(5)
  /*f:=1+2+3+4+5
  fmt.Println(f)*/
  fmt.Println(s)
}
//
func getSum(num int) int {

	if num==1 {
		return 1
	}
	return	getSum(num-1)+num
}
/*5+4
4+3
3+2
2+1
1*/