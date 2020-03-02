package main

import (
	"fmt"
)

type record struct {//structure delaration
	amount float32
	tax float32
	discount float32
}
var td, tt, tgs float32
func main()  {
	var r = []record{{10000, 10, 5},//compile time data input
		{10000, 1, 54},
		{2000, 80, 50},
		{123500, 50, 50},
		{21000, 10, 2},
		{470, 4, 8},
		{7832, 14, 8},
		{43787, 17, 15}}
	for _, rc := range r {//get each and every record one by one
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

func total(r record,tx,ds float32) {//total 
fmt.Println("Total Amount To Pay",r.amount,"+",tx,"=",r.amount+tx,"-",ds,"=",r.amount+tx-ds)
}

func grosssale(r record, tx, ds float32) float32 {//gross sale
	gross:=r.amount+tx-ds
	fmt.Println("Gross Sale for",r.amount,"->",gross)
	return gross
}

func netsale(r record,gs,ds float32) float32 {//net sale
	fmt.Println("Net Sale Amount ->",r.amount)
	return r.amount
}

func dis(r record, tx float32) float32 {//discount
	r.amount=r.amount+tx
	var disAmount = r.amount * r.discount / 100
	fmt.Println("Discount Amount for ",r.amount,"->",disAmount)
	return disAmount
}

func tax(r record) float32  {//tax
	taxAmount:= (r.amount)*(r.tax/100)
	fmt.Println("Tax Amount for ",r.amount,"->",taxAmount)
	return taxAmount
}
