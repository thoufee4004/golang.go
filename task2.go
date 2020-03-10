//OUTPUT
same as previous Program
but here it has insert
and calculation
and print code separately
its a initial try so may have some errror//


package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type users struct {
	Id    bson.ObjectId        `bson:"_id"`
	User       bson.ObjectId             `bson:"user"`
	Name string `bson:"name"`
	EmailId string `bson:"emailId"`
}
type sales struct {
	Id    bson.ObjectId        `bson:"_id"`
	NetSale    float32         `bson:"netSale"`
	Tax        float32         `bson:"tax"`
	Discount   float32         `bson:"discount"`
	User       bson.ObjectId             `bson:"user"`

}
type SalesCpy struct {
	Id    bson.ObjectId        `bson:"_id"`
	NetSale    float32         `bson:"netSale"`
	Tax        float32         `bson:"tax"`
	Discount   float32         `bson:"discount"`
	User       bson.ObjectId             `bson:"user"`


}

type FinalData struct {
	NetSale    float32
	Tax        float32
	Discount   float32

}
var U=[]users{
	{bson.NewObjectId(),bson.ObjectIdHex("5e621ce611bae6e8c8064b73"), "Tfk", "tfk@mail.com"},
	{bson.NewObjectId(), bson.ObjectIdHex("5e621ce611bae6e8c8064b74"),"msd", "hk@mail.com"},
	{bson.NewObjectId(), bson.ObjectIdHex("5e621ce611bae6e8c8064b75"),"gvm", "dd@mail.com"},
	{bson.NewObjectId(),bson.ObjectIdHex("5e621ce611bae6e8c8064b76"),"arr","arr@mail.com"},
}
var R = []sales{
	{bson.NewObjectId(), 10, 5,2, bson.ObjectIdHex("5e621ce611bae6e8c8064b73")},
	{bson.NewObjectId(), 100, 54 ,55, bson.ObjectIdHex("5e621ce611bae6e8c8064b74")},
	{bson.NewObjectId(), 8000, 50, 63, bson.ObjectIdHex("5e621ce611bae6e8c8064b74")},
	{bson.NewObjectId(), 5000, 50,7, bson.ObjectIdHex("5e621ce611bae6e8c8064b73")},
	{bson.NewObjectId(), 100, 2,8, bson.ObjectIdHex("5e621ce611bae6e8c8064b73")},
	{bson.NewObjectId(), 400, 8,4, bson.ObjectIdHex("5e621ce611bae6e8c8064b75")},
	{bson.NewObjectId(), 1400, 8,6, bson.ObjectIdHex("5e621ce611bae6e8c8064b75")},
	{bson.NewObjectId(), 17000, 15,3, bson.ObjectIdHex("5e621ce611bae6e8c8064b75")},
	{bson.NewObjectId(), 5444, 5,8, bson.ObjectIdHex("5e621ce611bae6e8c8064b73")},
	{bson.NewObjectId(),674464,6,8,bson.ObjectIdHex("5e621ce611bae6e8c8064b76")},
}

func main() {

	insert()
	//calculate()
	//UserData()
}



func insert() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
	con := session.DB("Task").C("sale")
	con2 := session.DB("Task").C("user")

	len, err := con.Count()
	if err != nil {
		return
	}
	log.Println("len :", len)

	fmt.Println("Connected to MongoDB!")

	for _, us := range U {
		_, in := con2.Upsert(us, us)
		if in != nil {
			fmt.Println(err)
		}
	}
	var res1 []users
	err1 := con2.Find(bson.M{}).All(&res1)
	if err1 != nil {
		return
	}
	fmt.Println("inserted: ", res1)

	for _, record := range R {
		_, in := con.Upsert(record, record)
		if in != nil {
			fmt.Println(err)
		}
	}
	var res []sales
	err = con.Find(bson.M{}).All(&res)
	if err != nil {
		return
	}
	fmt.Println("inserted: ", res)
}
/*var td, tt, tgs float32
func calculate() {

	for _, rc := range R {
		tx := tax(rc)
		ds := dis(rc, tx)
		gs := grosssale(rc, tx, ds)
		td += ds
		tt += tx
		tgs += gs
		fmt.Println("Gross:\n", gs,"\nTax:\n", tx,"\nDiscount:\n", ds)
	}
	fmt.Println("total:\n", td,"\nTotalDis:\n", tt,"\nTotalGross:\n", tgs)
}


func grosssale(R sales, tx, ds float32) float32 {
	gross:=R.NetSale+tx-ds
	fmt.Println("Gross Sale for",R.NetSale,"->",gross)
	return gross
}


func dis(R sales, tx float32) float32 {
	R.NetSale=R.NetSale+tx
	var disAmount = R.NetSale * R.Discount / 100
	fmt.Println("Discount Amount for ",R.NetSale,"->",disAmount)
	return disAmount
}

func tax(R sales) float32  {
	taxAmount:= (R.NetSale)*(R.Tax/100)
	fmt.Println("Tax Amount for ",R.NetSale,"->",taxAmount)
	return taxAmount
}*/

func UserData() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return
	}
	con := session.DB("Task").C("sale")
	con2 := session.DB("Task").C("user")
	Data := make(map[bson.ObjectId]FinalData)
	mapdata1 := make(map[bson.ObjectId]users)

	id_user := []bson.ObjectId{bson.ObjectIdHex("5e621ce611bae6e8c8064b73"), bson.ObjectIdHex("5e621ce611bae6e8c8064b74"), bson.ObjectIdHex("5e621ce611bae6e8c8064b75")}
	for _, id := range id_user {
		data := users{}
		//log.Println("user",id)
		query := bson.M{"name": 1, "user": 1}
		con2.Find(bson.M{"user": id}).Select(query).One(&data)
		//log.Println("data",data)
		mapdata1[id] = data
	}

	for _, v := range mapdata1 {
		fmt.Println("v :", v)
	}

	_ids := []bson.ObjectId{bson.ObjectIdHex("5e621ce611bae6e8c8064b73"), bson.ObjectIdHex("5e621ce611bae6e8c8064b74"), bson.ObjectIdHex("5e621ce611bae6e8c8064b75")}
	for _, userId := range _ids {
		var saleData []SalesCpy
		//log.Println("userId",userId)
		if err := con.Find(bson.M{"user": userId}).All(&saleData); err != nil {
			//log.Println
			return
		}
		//log.Println("saleData", saleData)
		finalData := FinalData{}
		for _, l := range saleData {
			finalData.Discount = finalData.Discount + l.Discount
			finalData.NetSale = finalData.NetSale + l.NetSale
			finalData.Tax = finalData.Tax + l.Tax
		}

		Data[userId] = finalData
	}
	fmt.Println("Total\n", Data)
	fmt.Println("CAlCULATIONS:\n")

}
