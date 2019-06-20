package main

import "fmt"
import "time"

func main(){

	i := 2
	switch i{

		case 1: 
			fmt.Println("I am 1")
		case 2: 
			fmt.Println("I am 2")
		case 3: 
			fmt.Println("I am 3")

	}

	switch time.Now().Weekday(){

		case time.Monday:
			fmt.Println("Today is Monday. ")
		case time.Tuesday:
			fmt.Println("Today is Tuesday. ")
		case time.Wednesday:
			fmt.Println("Today is Wednesday . ")
		case time.Thursday:
			fmt.Println("Today is Thursday. ")
		case time.Friday:
			fmt.Println("Today is Friday. ")
		case time.Saturday:
			fmt.Println("Today is Saturday. ")
		case time.Sunday:
			fmt.Println("Today is Sunday. ")

	}

}