/*package main
STRING REVERSE
import (
"fmt"
)

func main()  {
	var str ,str1 string
	str1 = ""
	fmt.Scanf("%s",&str)
	fmt.Println(str)
	for _, r:= range str{
		//defer fmt.Printf(string(r))
		str1=string(r)+str1
	}
	fmt.Println(str1)
}*/

/*package main
PATTERN PROGRAM
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print(i)
		}
		for j:=0;j<n-i-1;j++{
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
*/


/*package main
INSERTING A DATA TO LOCAL DB IN MONGODB
import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type Student struct {
	Name   string     `bson:"name"`
	Age    int        `bson:"age"`
	Mark   int        `bson:"mark"`

}

func main() {
	st:=[]Student{
		{"tfk", 22, 95},
		{"iron man",45,66},
		{"edith",2,100},
		{"tony",45,100},
		{"thor",55,90},
		{"hulk",55,98},
		{"cap",60,88},
	}

session,err:=mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
con:=session.DB("Task").C("students")

	len, err := con.Count()
	if err != nil {
		return
	}
	log.Println("len :",len)

	fmt.Println("Connected to MongoDB!")
	for _,record:= range st {
		in, _ :=con.Upsert(record,record)
		if in != nil {
			fmt.Println(err)
		}
	}
	var res []Student
	err = con.Find(bson.M{}).All(&res)
	if err != nil {
		return
	}
	//fmt.Println(res)
	fmt.Println("inserted: ", res)

}*/



/*package main
CALCULATION OF 
TAX,TOTAL TAX
DISCOUNT,TOTAL DISCOUNT
GROSS,TOTAL GROSS

import (
	"fmt"
)

type record struct {
	amount float32
	tax float32
	discount float32
}
var td, tt, tgs float32
func main()  {

	var r = []record{{10000, 10, 5},
		{10000, 1, 54},
		{2000, 80, 50},
		{123500, 50, 50},
		{21000, 10, 2},
		{470, 4, 8},
		{7832, 14, 8},
		{43787, 17, 15}}
	for _,v:= range r{
		fmt.Scanf("enter amount %f",&v.amount)
		fmt.Scanf("enter amount %f",&v.tax)
		fmt.Scanf("enter amount %f",&v.discount)

	}
	//fmt.Println("len",len(r))
	for _, rc := range r {
		tx := tax(rc)
		ds := dis(rc, tx)
		gs := grosssale(rc, tx, ds)
		netsale(rc, gs, ds)
		total(rc, tx, ds)
		td += ds
		tt += tx
		tgs += gs
		fmt.Println("Gross:\n", gs,"\nTax:\n", tx,"\nDiscount:\n", ds)}
	fmt.Println("total:\n", td,"\nTotalDis:\n", tt,"\nTotalGross:\n", tgs)
}

func total(r record,tx,ds float32) {
fmt.Println("Total Amount To Pay",r.amount,"+",tx,"=",r.amount+tx,"-",ds,"=",r.amount+tx-ds)
}

func grosssale(r record, tx, ds float32) float32 {
	gross:=r.amount+tx-ds
	fmt.Println("Gross Sale for",r.amount,"->",gross)
	return gross
}

func netsale(r record,gs,ds float32) float32 {
	fmt.Println("Net Sale Amount ->",r.amount)
	return r.amount
}

func dis(r record, tx float32) float32 {
	r.amount=r.amount+tx
	var disAmount = r.amount * r.discount / 100
	fmt.Println("Discount Amount for ",r.amount,"->",disAmount)
	return disAmount
}

func tax(r record) float32  {
	taxAmount:= (r.amount)*(r.tax/100)
	fmt.Println("Tax Amount for ",r.amount,"->",taxAmount)
	return taxAmount
}
*/

/*package main
JUST APPEND TWO SLICES
import "fmt"

func main()  {
	a,b:=[]int{1,2,3,4,5} ,[]int{6,7,8,9,10}
	fmt.Println("slice a ",a,"\nlength",len(a))
	a=append(a,b...)
	fmt.Println("after append",a,"\nlength",len(a))
}*/
