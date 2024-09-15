package main

import "fmt"

func main() {
	// while loop
	//while_loop()
	// for loop
	//for_loop()
	//infinite loop
	//infinite_loop()
	//for range
	for_range()	
}

func while_loop(){
	i:=1
	for i<=3{
		fmt.Println(i)
		i++
	}
}
func for_loop (){
for i:=0;i<=3;i++{
	//if i==2
		//continue
		//break
	
	fmt.Println(i)
	
}
}
func infinite_loop(){
	for{
		fmt.Println(1)
	}
}

func for_range(){
	for i := range 5{
		fmt.Println(i)
	}
}