package main

import "fmt"

func main(){



	var a [3] int;

	for j := 0; j < len(a); j++{
		a[j] = j+1;
		fmt.Println(a[j]);
	}	

	b := [5] int{1, 2, 3, 4, 5}


	for j := 0; j < len(b); j++{
		fmt.Println(b[j]);
	}	

	var d [5] int 

	for j := 0; j < len(a); j++{
		d[j] = j+1;
		fmt.Println("this is", d[j]);
	}

	var c [5][5] int
	for j := 0; j < 5; j++{
		for z := 0; z < 5; z++{
			c[j][z] = z+j;
		}
	}
	fmt.Println(c);


}