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
}
func newOrder(id,status string,amount float32)*Order{

	myOrder := Order{
		id: id,
	   status: status,
	   amount: amount,
   }
   return &myOrder

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
	myOder := newOrder("1","revived",20)
	fmt.Println(myOder);
}