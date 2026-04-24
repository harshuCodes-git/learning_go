package main

import (
	"fmt"
	"time"
)

func main() {
	age := 18
	switch age {
	case 18:
		println("Is an adult")
	case 17:
		println("Not an adult")
	default:
		println("Invalid age")
	}

	// multi condition switch
	times := time.Now().Weekday()
	fmt.Println(times);
	switch times{
	case time.Saturday, time.Sunday:
		println("It's a weekend")
	default:
		println("It's a weekday")
	}

	// type switch 

	whoAmI:= func(i any){
		switch v:=i.(type){
		case int:
			fmt.Printf("I am an integer and my value is %v\n", v)
		case string:
			fmt.Printf("I am a string and my value is %v\n", v)
		default:
			fmt.Printf("I am of a different type and my value is %v\n", v)
		}
	}
	whoAmI(42)
	whoAmI("Hello, Go!")
	whoAmI(3.14)		

}