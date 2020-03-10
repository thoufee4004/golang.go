/*OUTPUT
Tfk
Total Net Sale-> 10554
total Tax -> 2774.7000122070312
Total Discount -> 793.719989001751
Total Gross 12534.98002320528
msd
Total Net Sale-> 8100
total Tax -> 4054.0000038146973
Total Discount -> 5095
Total Gross 7059.000003814697
gvm
Total Net Sale-> 18800
total Tax -> 2694
Total Discount -> 610
Total Gross 20884
arr
Total Net Sale-> 674464
total Tax -> 40467.83984375
Total Discount -> 53957.1171875
Total Gross 660974.72265625

Process finished with exit code 0
*/

package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	//"log"
)

type users struct {
	Id      bson.ObjectId `bson:"_id"`
	User    bson.ObjectId `bson:"user"`
	Name    string        `bson:"name"`
	EmailId string        `bson:"emailId"`
}

type sales struct {
	Id    bson.ObjectId        `bson:"_id"`
	NetSale    float32         `bson:"netSale"`
	Tax        float32         `bson:"tax"`
	Discount   float32         `bson:"discount"`
	User       bson.ObjectId             `bson:"user"`

}

type Final struct {
	NetSale    float32         `bson:"netSale"`
	Tax        float32         `bson:"tax"`
	Discount   float32         `bson:"discount"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	connection1 := session.DB("Task").C("sale")
	connection2 := session.DB("Task").C("user")
	data1 := []users{}
	data2 := []sales{}
	user := connection2.Find(bson.M{}).All(&data1)
	if user != nil {
		return
	}
	//finalData := make(map[string]interface{})
	for _, v := range data1 {
		//name := v.Name
		//log.Println("No.of.Rec", len(data2))
		sale := connection1.Find(bson.M{"user": v.User}).All(&data2)
		if sale != nil {
			return
		}
		netSaleTotal := 0.0
		discountTotal := 0.0
		taxTotal := 0.0
		grossTotal:=0.0
		for _, l := range data2 {
			netSaleTotal = netSaleTotal + float64(l.NetSale)
			discountTotal = discountTotal + float64(l.NetSale*(l.Discount/100))
			taxTotal = taxTotal + float64(l.NetSale*(l.Tax/100))
			grossTotal=netSaleTotal+taxTotal-discountTotal
			}
		fmt.Println(v.Name)
		fmt.Println("Total Net Sale->",netSaleTotal)
		fmt.Println("total Tax ->",taxTotal)
		fmt.Println("Total Discount ->",discountTotal)
		fmt.Println("Total Gross",grossTotal)

			//finalData[v.Name"/"+ "Net Sale"] = netSaleTotal
			//finalData[v.Name+"/"+"Discount"] = discountTotal
			//finalData[v.Name+ "/" +"Tax"] = taxTotal
	}
	//fmt.Println("\nfinalData",finalData)
}
