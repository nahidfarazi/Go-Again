package main

import (
	"fmt"
	"time"
)

//order struct
type Order struct{
	id 			string
	amount 		float32
	status 		string
	createdAt 	time.Time
	Customer
}
// customer struct
type Customer struct{
	name string
	phone string
}
func newOrder(id,status string,amount float32)*Order{

	myOrder := Order{
		id: id,
	   status: status,
	   amount: amount,
   }
   return &myOrder

}

func newCustomer(name, phone string)*Customer{
	newCustomer:=Customer{
		name: name,
		phone: phone,
	}
	return &newCustomer
}
//receiver type
func (o *Order) changeStatus(status string){
	o.status = status

}
func (o Order) getAmount()float32{
	return o.amount
}
func (o *Order) changeId(id string){
	o.id = id

}
func main() {
	
	
	// myOrder.changeStatus("confirmed")
	// myOrder.changeId("1")
	// // myOrder.createdAt =time.Now()
	// // fmt.Println(myOrder);
	// fmt.Println(myOrder);
	// fmt.Println(myOrder.getAmount());
	
	myOder := newOrder("1","revived",20,)
	myCustom := newCustomer("Nahid Farazi","01795543090")
	fmt.Println(myOder);
	fmt.Println(myCustom);
	
	
}