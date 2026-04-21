package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i); 
		i++
	}

	// infinite loop 
	// for {
	// 	fmt.Println("harsh dubey")
	// }

	for init:=0; init<=3; init++{
		fmt.Println(init)
	}

	// for range loop
	for val:=range 10 {
		fmt.Println(val)
		
	}
}