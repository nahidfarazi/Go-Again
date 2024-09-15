package main

import "fmt"

func main() {
	createMap()
}

func createMap(){
	/*m := make(map[string]string)
	// set element
	m["name"] = "golang"
	m["company"] = "google.com"

	// get element
	fmt.Println(m["name"])
	fmt.Println(m["company"]);*/

	m := map[string]int{"language":1,"age":20}
	//fmt.Println(m);
	_, ok :=m["age "]
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}
}