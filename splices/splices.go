package main 

import "fmt"

func main(){

	s := make([]int, 3, 4)
	fmt.Println(s)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(s)

	for j := 0; j <= 5; j++{
		s = append(s, j)
	}

	fmt.Println(s)

	t := make([]int, 8)

	copy(t, s)

	fmt.Println(t)

	c := t[:5]

	fmt.Println(c)

	var x *int 
	x = nil

	var z int 
	z = nil
	
}