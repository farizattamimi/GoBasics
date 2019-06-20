package main

import "fmt"

func add(a int, b int)int{
	return a+b
}

func mult(a int, b int)int{
	return a*b
}

func sub(a int, b int)int{
	return a-b
}

func div(a int, b int)int{
	return a/b
}

func main(){
	
	res := add(5, 4)
	fmt.Println(res)

	res = mult(5, 4)
	fmt.Println(res)

	res = sub(5, 4)
	fmt.Println(res)

	res = div(5, 4)
	fmt.Println(res)
}