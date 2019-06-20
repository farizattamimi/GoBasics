package main

import "encoding/json"
import "fmt"

type person struct{

	ID int
	Name string
	Color [] string
	Age int
	Youth bool

}

func main(){

person1 := person{
	ID: 1,
	Name: "Fariz",
	Color: [] string {"blue", "red", "green"},
	Age: 22,
	Youth: false}
person1temp,_ := json.Marshal(&person1)
fmt.Println(string(person1temp))


str := `{"ID": 2, "Name": "John", "Color": ["Gold", "Silver"], "Age": 22, "Youth": false}`
var person2 person
json.Unmarshal([]byte(str), &person2)
fmt.Println(person2)

}