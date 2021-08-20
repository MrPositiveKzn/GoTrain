package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int, c chan int){
	time.Sleep(3 *time.Second)
	fmt.Println("zzzzz", id)
	c <- id
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++{
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++{
		gopherID := <- c
		fmt.Println("gopher", gopherID, "finished sleep")
	}
	time.Sleep(4 *time.Second)
}

