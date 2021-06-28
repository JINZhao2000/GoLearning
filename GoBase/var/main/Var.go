package main

import (
	"fmt"
	"strconv"
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
	fmt.Printf("Type: %T,Bytes: %d\n", A, unsafe.Sizeof(A))

	var char1 byte = ' '
	var char2 byte = '0'
	var char3 byte = 'a'

	fmt.Printf("%c %c %c\n", char1, char2, char3)

	var char4 rune = '中'
	fmt.Printf("%c\n", char4)

	for i := 10000; i < 10050; i++ {
		fmt.Printf("%c\n",i)
	}

	var str1 = `str\nstr`
	fmt.Println(str1)

	var str2 = `import (
	"fmt"
	"unsafe"
)`
	fmt.Println(str2)

	var str3 = fmt.Sprintf("%c%c%c",char2, char3, char4)
	fmt.Println(str3)
	str3 = fmt.Sprintf("%t",c)
	fmt.Println(str3)
	fmt.Printf("%q",str3)
	str3 = strconv.FormatInt(20,10);
	fmt.Println(str3)
	str3 = strconv.FormatFloat(1.2345678,'f',3,64)
	fmt.Println(str3)
	str3 = strconv.FormatBool(true)
	fmt.Println(str3)
}
