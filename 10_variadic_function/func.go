package main

import "fmt"

func main() {
	mySlice:=[] int {2,3,4,1}
	 result := sum(mySlice...)
	 fmt.Println(result);
}
func sum(num ...int)int{
	sum :=0
	for _,i :=range num{
		sum+=i
	}
	return sum
}