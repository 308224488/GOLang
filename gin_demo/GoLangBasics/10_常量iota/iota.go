package main

import (
	"fmt"
)

func main()  {
	const(
		a=iota
		b
		c="李遇"
		d=iota
		e
		)
	const (
		f=iota
		g
		h=false
	)
	fmt.Println(a,b,c,d,e,f,g,h)
}