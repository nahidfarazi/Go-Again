package main

import "fmt"

func main() {
	// num := []int{4,5,8}
	// for i, nums := range num{
	// 	fmt.Println(i," - ",nums)

	// }
	mapRange()
}

func mapRange(){
	m:= map[string] string{"first_Name":"Nahid","last_Name":"Farazi"}

	for k,v:= range m{
		fmt.Println(k,":",v);
	}
}