package main

import "fmt"

func display1(c int){
	fmt.Println("from display1")
	fmt.Println(c)
}
func display2(c int){
	fmt.Println("from display2")
	fmt.Println(c)
}
func add(a [] int,c chan int) {

	sum:=0
	for _,i:=range a{
		sum+=i
	}
	c<-sum
	fmt.Println("here",)
}
func main(){
	a:=[]int{12,-4,57,89,-63,-254,12,58,93,61,-44,-54}
	c:=make(chan int)
	n:=len(a)/2
	display1(6)
	display2(8)
	go add(a[:n],c)
	x:=<-c
	go add(a[n:],c)
	y:=<-c
	fmt.Println("'len/2:'",x,"\n","':len/2'",y,"\n","add",x+y)
}

//output
from display1
6
from display2
8
here
here
'len/2:' -163 
 ':len/2' 126 
 add -37
