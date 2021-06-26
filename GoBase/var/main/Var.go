package main

import (
	"fmt"
	"unsafe"
)

var (
	A = 100
	B = 200
)

func main()  {
	var a int = 0
	var b = "1"
	c := true
	c = false
	fmt.Println(a,b,c)
	var d, e, f = 1,"1",1.0
	fmt.Println(d,e,f)
	fmt.Println(A, B)
	fmt.Printf("Type: %T,Bytes: %d", A, unsafe.Sizeof(A));
}
