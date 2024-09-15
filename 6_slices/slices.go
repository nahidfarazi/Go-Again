package main

import "fmt"

func main() {
	
/*var num = make([]int,0,5)
num = append(num, 40)
num = append(num, 58)
num = append(num, 32)
num = append(num, 1000)
num = append(num, 35)
num = append(num, 265)
fmt.Println(num);
fmt.Println(cap(num));
fmt.Println(len(num));*/
//capArray()
sliceOp()
}

func capArray(){
	num := [] int {}
	num = append(num, 2)
	var nam2 = make([]int, len(num))
// 	num = append(num, 10)
// 	num = append(num, 54)
// 	num = append(num, 54)
// 	fmt.Println(num);
// 	fmt.Println(cap(num));
// 	fmt.Println(len(num));
copy(nam2, num)
fmt.Println(num, nam2);
}

func sliceOp(){
	var num = []int{1,2,3}
	fmt.Println(num[:2])
}