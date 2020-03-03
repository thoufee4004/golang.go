package main

import "fmt"

func main()  {
	a,b:=[]int{1,2,3,4,5} ,[]int{6,7,8,9,10}
	fmt.Println("slice a ",a,"\nlength",len(a))
	a=append(a,b...)
	fmt.Println("after append",a,"\nlength",len(a))
}
