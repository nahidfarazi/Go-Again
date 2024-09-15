package main

import (
	"fmt"
	"time"
)

func main() {
	  //simpleSwitch()
	  //multiSwitch()
	  typeSwitch()
}

func simpleSwitch(){
	i := 7
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("tow")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("out of range")
	}
}

func multiSwitch(){
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work day")
	}
}

func typeSwitch(){
	whoAmI := func (i interface{})  {
		switch t := i.(type){
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("other",t)
		}
	}
	whoAmI(10.4)
}