package main

import "fmt"

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment());
}
func counter()func()int{
	count :=0
	return func()int{
		count+=1
		return count
	}
}