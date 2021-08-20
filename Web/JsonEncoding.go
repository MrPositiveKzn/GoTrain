package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//type User struct {
//	id			int
//	name		string
//	occupation	string
//}

func main() {
	u1 := User{id: 1, name: "Kolya", occupation: "Backend developer"}
	json_data, err := json.Marshal(u1)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(json_data)
	users := []User{
		{id: 2, name: "Vasya", occupation: "Java dev"},
		{id: 3, name: "Vlad", occupation: "Go dev"},
	}
	json_data2, err := json.Marshal(users)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(json_data2)
}
