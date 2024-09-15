package main

func main() {
	fn := func (a int) int {
		return 2
	}
	processIt(fn)
	// ADD_result := add(10,22)
	// fmt.Println(ADD_result);
	// SUB_result := sub(34,22)
	// fmt.Println(SUB_result);
	// get_value1,get_value2,get_value3 := getLang()
	// fmt.Println(get_value1,get_value2,get_value3);

}
func getLang() (string,string,string){
	return "golang","java","c"
}

func add(a int,b int) int{
	return a+b
}
func sub(x,y int)int{
	return x-y
}

func processIt(fn func(a int)int){
	fn(10)
}