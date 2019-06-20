package main

import "fmt"

func main(){
	
	maps := make(map[string]int)

	maps["fariz"] = 1
	maps["jack"] = 2
	maps["john"] = 3

	fmt.Println(maps)

	name := maps["fariz"]

	fmt.Println(name);

	fmt.Println(len(maps))

	delete(maps, "jack")

	fmt.Println(maps)

	p, there := maps["fariz"]

	fmt.Println(p)
	fmt.Println(there)

	p, there = maps["jack"]

	fmt.Println(p)
	fmt.Println(there)

	 n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}