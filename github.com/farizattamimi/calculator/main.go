package main

import "fmt"
import "strconv"
import "github.com/farizattamimi/calculator/math"
import "os"
import "bufio"

func main (){

	c := math.NewCalculator()

	fmt.Println("--------------------------------")
	fmt.Println(".       Calculator App          ")
	fmt.Println("--------------------------------")
	fmt.Println()

		fmt.Println("1. Add")
		fmt.Println("2. Substract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Div")
		fmt.Println("0. Exit")

		

	for{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		selection := scanner.Text()

		if selection == "1"{
			fmt.Println("Insert your first option:")
			scanner.Scan()
			firstnum := scanner.Text()
			fmt.Println("Insert your second option:")
			scanner.Scan()
			secondnum := scanner.Text()

			one, err := strconv.Atoi(firstnum)
			if err != nil {
				panic(err)
			}
			two, err := strconv.Atoi(secondnum)
			if err != nil {
				panic(err)
			}
			//c.SetFirstNumber(one)
			//c.SetSecondNumber(two)
			result := c.Add(one, two)

			fmt.Println("Result: ", one, " + ", two, " = ", result)

		}else if selection == "2"{
			fmt.Println("Insert your first option:")
			scanner.Scan()
			firstnum := scanner.Text()
			fmt.Println("Insert your second option:")
			scanner.Scan()
			secondnum := scanner.Text()

			one, err := strconv.Atoi(firstnum)
			if err != nil {
				panic(err)
			}
			two, err := strconv.Atoi(secondnum)
			if err != nil {
				panic(err)
			}
			result := c.Sub(one, two)

			fmt.Println("Result: ", one, " - ", two, " = ", result)

		}else if selection == "3"{
			fmt.Println("Insert your first option:")
			scanner.Scan()
			firstnum := scanner.Text()
			fmt.Println("Insert your second option:")
			scanner.Scan()
			secondnum := scanner.Text()

			one, err := strconv.Atoi(firstnum)
			if err != nil {
				panic(err)
			}
			two, err := strconv.Atoi(secondnum)
			if err != nil {
				panic(err)
			}
			result := c.Mult(one, two)

			fmt.Println("Result: ", one, " * ", two, " = ", result)
			
		}else if selection == "4"{
			fmt.Println("Insert your first option:")
			scanner.Scan()
			firstnum := scanner.Text()
			fmt.Println("Insert your second option:")
			scanner.Scan()
			secondnum := scanner.Text()

			one, err := strconv.Atoi(firstnum)
			if err != nil {
				panic(err)
			}
			two, err := strconv.Atoi(secondnum)
			if err != nil {
				panic(err)
			}
			result, err := c.Div(one, two)

			if err == nil{
				fmt.Println("Result: ", one, " / ", two, " = ", result)
			}else{
				fmt.Println("Result: ", one, " / ", two, " = ", "undefined")
			}
			
		}else if selection == "0" {
			break;
		}else{
			for{
				fmt.Println("Try again. Please enter a number from 0 to 4")
				fmt.Println("1. Add")
				fmt.Println("2. Substract")
				fmt.Println("3. Multiply")
				fmt.Println("4. Div")
				fmt.Println("0. Exit")
				scanner.Scan()
				selection = scanner.Text()
				if(selection == "1" || selection == "2" ||selection == "3" ||selection == "4" ||selection == "0"){
					break;
				}
			}
			continue;
		}
		fmt.Println("1. Add")
		fmt.Println("2. Substract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Div")
		fmt.Println("0. Exit")

	}

	fmt.Println("Thank you. You have exited.")


}