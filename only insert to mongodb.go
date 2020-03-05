package main

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

}
